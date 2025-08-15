package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/samber/lo"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
)

func main() {
	fmt.Println("Downloading vanilla items NBT file...")
	resp, err := http.Get("https://github.com/df-mc/dragonfly/raw/refs/heads/master/server/world/vanilla_items.nbt")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("failed to download vanilla items NBT file: %s", resp.Status))
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("failed to read vanilla items NBT file: %w", err))
	}

	var m map[string]struct {
		RuntimeID      int32          `nbt:"runtime_id"`
		ComponentBased bool           `nbt:"component_based"`
		Version        int32          `nbt:"version"`
		Data           map[string]any `nbt:"data,omitempty"`
	}
	if err := nbt.Unmarshal(bytes, &m); err != nil {
		panic(fmt.Errorf("failed to unmarshal vanilla items NBT file: %w", err))
	}

	type itemEntry struct {
		RuntimeID      int16  `json:"runtime_id"`
		ComponentBased bool   `json:"component_based"`
		Version        *uint8 `json:"version"`
		Data           string `json:"component_nbt,omitempty"`
	}

	m2 := make(map[string]itemEntry)
	defData := make(map[string]any)
	for k, v := range m {
		var data string
		if v.Data == nil {
			v.Data = defData
		}
		dataBytes, err := nbt.Marshal(v.Data)
		if err != nil {
			panic(fmt.Errorf("failed to marshal component NBT for item %s: %w", k, err))
		}
		data = base64.StdEncoding.EncodeToString(dataBytes)
		m2[k] = itemEntry{
			RuntimeID:      int16(v.RuntimeID),
			ComponentBased: v.ComponentBased,
			Version:        lo.ToPtr(uint8(v.Version)),
			Data:           data,
		}
	}
	jsonData, err := json.MarshalIndent(m2, "", "    ")
	if err != nil {
		panic(fmt.Errorf("failed to marshal item data to JSON: %w", err))
	}
	fmt.Println("Writing item data to dragonfly_items.json...")
	if err := os.WriteFile("dragonfly_items.json", jsonData, 0644); err != nil {
		panic(fmt.Errorf("failed to write item data to dragonfly_items.json: %w", err))
	}
	fmt.Println("Item data written to dragonfly_items.json successfully.")
}
