package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func LevelSoundEvent(io protocol.IO, pk *packet.LevelSoundEvent) {
	if proto.IsProtoGTE(io, proto.ID1001) {
		soundType := soundEventToString(pk.SoundType)
		io.String(&soundType)
		soundEventFromString(io, &pk.SoundType, soundType)
	} else {
		io.Varuint32(&pk.SoundType)
	}
	io.Vec3(&pk.Position)
	io.Varint32(&pk.ExtraData)
	io.String(&pk.EntityType)
	io.Bool(&pk.BabyMob)
	io.Bool(&pk.DisableRelativeVolume)
	if proto.IsProtoGTE(io, proto.ID786) {
		io.Int64(&pk.EntityUniqueID)
		if proto.IsProtoGTE(io, proto.ID975) {
			protocol.OptionalFunc(io, &pk.FireAtPosition, io.Vec3)
		}
	}
}

func soundEventToString(x uint32) string {
	switch x {
	case 0:
		return "item_use_on"
	case 1:
		return "hit"
	case 2:
		return "step"
	case 3:
		return "fly"
	case 4:
		return "jump"
	case 5:
		return "break"
	case 6:
		return "place"
	case 7:
		return "heavy_step"
	case 8:
		return "gallop"
	case 9:
		return "fall"
	case 10:
		return "ambient"
	case 11:
		return "ambient_baby"
	case 12:
		return "ambient_in_water"
	case 13:
		return "breathe"
	case 14:
		return "death"
	case 15:
		return "death_in_water"
	case 16:
		return "death_to_zombie"
	case 17:
		return "hurt"
	case 18:
		return "hurt_in_water"
	case 19:
		return "mad"
	case 20:
		return "boost"
	case 21:
		return "bow"
	case 22:
		return "squish_big"
	case 23:
		return "squish_small"
	case 24:
		return "fall_big"
	case 25:
		return "fall_small"
	case 26:
		return "splash"
	case 27:
		return "fizz"
	case 28:
		return "flap"
	case 29:
		return "swim"
	case 30:
		return "drink"
	case 31:
		return "eat"
	case 32:
		return "takeoff"
	case 33:
		return "shake"
	case 34:
		return "plop"
	case 35:
		return "land"
	case 36:
		return "saddle"
	case 37:
		return "armor"
	case 38:
		return "mob_armor_stand_place"
	case 39:
		return "add_chest"
	case 40:
		return "throw"
	case 41:
		return "attack"
	case 42:
		return "attack_nodamage"
	case 43:
		return "attack_strong"
	case 44:
		return "warn"
	case 45:
		return "shear"
	case 46:
		return "milk"
	case 47:
		return "thunder"
	case 48:
		return "explode"
	case 49:
		return "fire"
	case 50:
		return "ignite"
	case 51:
		return "fuse"
	case 52:
		return "stare"
	case 53:
		return "spawn"
	case 54:
		return "shoot"
	case 55:
		return "break_block"
	case 56:
		return "launch"
	case 57:
		return "blast"
	case 58:
		return "large_blast"
	case 59:
		return "twinkle"
	case 60:
		return "remedy"
	case 61:
		return "unfect"
	case 62:
		return "levelup"
	case 63:
		return "bow_hit"
	case 64:
		return "bullet_hit"
	case 65:
		return "extinguish_fire"
	case 66:
		return "item_fizz"
	case 67:
		return "chest_open"
	case 68:
		return "chest_closed"
	case 69:
		return "shulkerbox_open"
	case 70:
		return "shulkerbox_closed"
	case 71:
		return "enderchest_open"
	case 72:
		return "enderchest_closed"
	case 73:
		return "power_on"
	case 74:
		return "power_off"
	case 75:
		return "attach"
	case 76:
		return "detach"
	case 77:
		return "deny"
	case 78:
		return "tripod"
	case 79:
		return "pop"
	case 80:
		return "drop_slot"
	case 81:
		return "note"
	case 82:
		return "thorns"
	case 83:
		return "piston_in"
	case 84:
		return "piston_out"
	case 85:
		return "portal"
	case 86:
		return "water"
	case 87:
		return "lava_pop"
	case 88:
		return "lava"
	case 89:
		return "burp"
	case 90:
		return "bucket_fill_water"
	case 91:
		return "bucket_fill_lava"
	case 92:
		return "bucket_empty_water"
	case 93:
		return "bucket_empty_lava"
	case 94:
		return "armor_equip_chain"
	case 95:
		return "armor_equip_diamond"
	case 96:
		return "armor_equip_generic"
	case 97:
		return "armor_equip_gold"
	case 98:
		return "armor_equip_iron"
	case 99:
		return "armor_equip_leather"
	case 100:
		return "armor_equip_elytra"
	case 101:
		return "record_13"
	case 102:
		return "record_cat"
	case 103:
		return "record_blocks"
	case 104:
		return "record_chirp"
	case 105:
		return "record_far"
	case 106:
		return "record_mall"
	case 107:
		return "record_mellohi"
	case 108:
		return "record_stal"
	case 109:
		return "record_strad"
	case 110:
		return "record_ward"
	case 111:
		return "record_11"
	case 112:
		return "record_wait"
	case 113:
		return "stop_record"
	case 114:
		return "flop"
	case 115:
		return "elderguardian_curse"
	case 116:
		return "mob_warning"
	case 117:
		return "mob_warning_baby"
	case 118:
		return "teleport"
	case 119:
		return "shulker_open"
	case 120:
		return "shulker_close"
	case 121:
		return "haggle"
	case 122:
		return "haggle_yes"
	case 123:
		return "haggle_no"
	case 124:
		return "haggle_idle"
	case 125:
		return "chorusgrow"
	case 126:
		return "chorusdeath"
	case 127:
		return "glass"
	case 128:
		return "potion_brewed"
	case 129:
		return "cast_spell"
	case 130:
		return "prepare_attack"
	case 131:
		return "prepare_summon"
	case 132:
		return "prepare_wololo"
	case 133:
		return "fang"
	case 134:
		return "charge"
	case 135:
		return "camera_take_picture"
	case 136:
		return "leashknot_place"
	case 137:
		return "leashknot_break"
	case 138:
		return "growl"
	case 139:
		return "whine"
	case 140:
		return "pant"
	case 141:
		return "purr"
	case 142:
		return "purreow"
	case 143:
		return "death_min_volume"
	case 144:
		return "death_mid_volume"
	case 145:
		return "imitate_blaze"
	case 146:
		return "imitate_cave_spider"
	case 147:
		return "imitate_creeper"
	case 148:
		return "imitate_elder_guardian"
	case 149:
		return "imitate_ender_dragon"
	case 150:
		return "imitate_enderman"
	case 151:
		return "imitate_endermite"
	case 152:
		return "imitate_evocation_illager"
	case 153:
		return "imitate_ghast"
	case 154:
		return "imitate_husk"
	case 156:
		return "imitate_magma_cube"
	case 157:
		return "imitate_polar_bear"
	case 158:
		return "imitate_shulker"
	case 159:
		return "imitate_silverfish"
	case 160:
		return "imitate_skeleton"
	case 161:
		return "imitate_slime"
	case 162:
		return "imitate_spider"
	case 163:
		return "imitate_stray"
	case 164:
		return "imitate_vex"
	case 165:
		return "imitate_vindication_illager"
	case 166:
		return "imitate_witch"
	case 167:
		return "imitate_wither"
	case 168:
		return "imitate_wither_skeleton"
	case 169:
		return "imitate_wolf"
	case 170:
		return "imitate_zombie"
	case 171:
		return "imitate_zombie_pigman"
	case 172:
		return "imitate_zombie_villager"
	case 173:
		return "block_end_portal_frame_fill"
	case 174:
		return "block_end_portal_spawn"
	case 175:
		return "random_anvil_use"
	case 176:
		return "bottle_dragonbreath"
	case 177:
		return "portal_travel"
	case 178:
		return "item_trident_hit"
	case 179:
		return "item_trident_return"
	case 180:
		return "item_trident_riptide_1"
	case 181:
		return "item_trident_riptide_2"
	case 182:
		return "item_trident_riptide_3"
	case 183:
		return "item_trident_throw"
	case 184:
		return "item_trident_thunder"
	case 185:
		return "item_trident_hit_ground"
	case 186:
		return "default"
	case 187:
		return "block_fletching_table_use"
	case 188:
		return "elemconstruct_open"
	case 189:
		return "icebomb_hit"
	case 190:
		return "balloonpop"
	case 191:
		return "lt_reaction_icebomb"
	case 192:
		return "lt_reaction_bleach"
	case 193:
		return "lt_reaction_epaste"
	case 194:
		return "lt_reaction_epaste2"
	case 195:
		return "lt_reaction_glow_stick"
	case 196:
		return "lt_reaction_glow_stick_2"
	case 197:
		return "lt_reaction_luminol"
	case 198:
		return "lt_reaction_salt"
	case 199:
		return "lt_reaction_fertilizer"
	case 200:
		return "lt_reaction_fireball"
	case 201:
		return "lt_reaction_mgsalt"
	case 202:
		return "lt_reaction_miscfire"
	case 203:
		return "lt_reaction_fire"
	case 204:
		return "lt_reaction_miscexplosion"
	case 205:
		return "lt_reaction_miscmystical"
	case 206:
		return "lt_reaction_miscmystical2"
	case 207:
		return "lt_reaction_product"
	case 208:
		return "sparkler_use"
	case 209:
		return "glowstick_use"
	case 210:
		return "sparkler_active"
	case 211:
		return "convert_to_drowned"
	case 212:
		return "bucket_fill_fish"
	case 213:
		return "bucket_empty_fish"
	case 214:
		return "bubble_up"
	case 215:
		return "bubble_down"
	case 216:
		return "bubble_pop"
	case 217:
		return "bubble_upinside"
	case 218:
		return "bubble_downinside"
	case 219:
		return "hurt_baby"
	case 220:
		return "death_baby"
	case 221:
		return "step_baby"
	case 222:
		return "spawn_baby"
	case 223:
		return "born"
	case 224:
		return "block_turtle_egg_break"
	case 225:
		return "block_turtle_egg_crack"
	case 226:
		return "block_turtle_egg_hatch"
	case 227:
		return "lay_egg"
	case 228:
		return "block_turtle_egg_attack"
	case 229:
		return "beacon_activate"
	case 230:
		return "beacon_ambient"
	case 231:
		return "beacon_deactivate"
	case 232:
		return "beacon_power"
	case 233:
		return "conduit_activate"
	case 234:
		return "conduit_ambient"
	case 235:
		return "conduit_attack"
	case 236:
		return "conduit_deactivate"
	case 237:
		return "conduit_short"
	case 238:
		return "swoop"
	case 239:
		return "block_bamboo_sapling_place"
	case 240:
		return "presneeze"
	case 241:
		return "sneeze"
	case 242:
		return "ambient_tame"
	case 243:
		return "scared"
	case 244:
		return "block_scaffolding_climb"
	case 245:
		return "crossbow_loading_start"
	case 246:
		return "crossbow_loading_middle"
	case 247:
		return "crossbow_loading_end"
	case 248:
		return "crossbow_shoot"
	case 249:
		return "crossbow_quick_charge_start"
	case 250:
		return "crossbow_quick_charge_middle"
	case 251:
		return "crossbow_quick_charge_end"
	case 252:
		return "ambient_aggressive"
	case 253:
		return "ambient_worried"
	case 254:
		return "cant_breed"
	case 255:
		return "item_shield_block"
	case 256:
		return "item_book_put"
	case 257:
		return "block_grindstone_use"
	case 258:
		return "block_bell_hit"
	case 259:
		return "block_campfire_crackle"
	case 260:
		return "roar"
	case 261:
		return "stun"
	case 262:
		return "block_sweet_berry_bush_hurt"
	case 263:
		return "block_sweet_berry_bush_pick"
	case 264:
		return "block_cartography_table_use"
	case 265:
		return "block_stonecutter_use"
	case 266:
		return "block_composter_empty"
	case 267:
		return "block_composter_fill"
	case 268:
		return "block_composter_fill_success"
	case 269:
		return "block_composter_ready"
	case 270:
		return "block_barrel_open"
	case 271:
		return "block_barrel_close"
	case 272:
		return "raid_horn"
	case 273:
		return "block_loom_use"
	case 274:
		return "ambient_in_raid"
	case 275:
		return "ui_cartography_table_take_result"
	case 276:
		return "ui_stonecutter_take_result"
	case 277:
		return "ui_loom_take_result"
	case 278:
		return "block_smoker_smoke"
	case 279:
		return "block_blastfurnace_fire_crackle"
	case 280:
		return "block_smithing_table_use"
	case 281:
		return "screech"
	case 282:
		return "sleep"
	case 283:
		return "block_furnace_lit"
	case 284:
		return "convert_mooshroom"
	case 285:
		return "milk_suspiciously"
	case 286:
		return "celebrate"
	case 287:
		return "jump_prevent"
	case 288:
		return "ambient_pollinate"
	case 289:
		return "block_beehive_drip"
	case 290:
		return "block_beehive_enter"
	case 291:
		return "block_beehive_exit"
	case 292:
		return "block_beehive_work"
	case 293:
		return "block_beehive_shear"
	case 294:
		return "drink_honey"
	case 295:
		return "ambient_cave"
	case 296:
		return "retreat"
	case 297:
		return "converted_to_zombified"
	case 298:
		return "admire"
	case 299:
		return "step_lava"
	case 300:
		return "tempt"
	case 301:
		return "panic"
	case 302:
		return "angry"
	case 303:
		return "ambient_warped_forest_mood"
	case 304:
		return "ambient_soulsand_valley_mood"
	case 305:
		return "ambient_nether_wastes_mood"
	case 306:
		return "ambient_basalt_deltas_mood"
	case 307:
		return "ambient_crimson_forest_mood"
	case 308:
		return "respawn_anchor_charge"
	case 309:
		return "respawn_anchor_deplete"
	case 310:
		return "respawn_anchor_set_spawn"
	case 311:
		return "respawn_anchor_ambient"
	case 312:
		return "particle_soul_escape_quiet"
	case 313:
		return "particle_soul_escape_loud"
	case 314:
		return "record_pigstep"
	case 315:
		return "lodestone_compass_link_compass_to_lodestone"
	case 316:
		return "smithing_table_use"
	case 317:
		return "armor_equip_netherite"
	case 318:
		return "ambient_warped_forest_loop"
	case 319:
		return "ambient_soulsand_valley_loop"
	case 320:
		return "ambient_nether_wastes_loop"
	case 321:
		return "ambient_basalt_deltas_loop"
	case 322:
		return "ambient_crimson_forest_loop"
	case 323:
		return "ambient_warped_forest_additions"
	case 324:
		return "ambient_soulsand_valley_additions"
	case 325:
		return "ambient_nether_wastes_additions"
	case 326:
		return "ambient_basalt_deltas_additions"
	case 327:
		return "ambient_crimson_forest_additions"
	case 328:
		return "power_on_sculk_sensor"
	case 329:
		return "power_off_sculk_sensor"
	case 330:
		return "bucket_fill_powder_snow"
	case 331:
		return "bucket_empty_powder_snow"
	case 332:
		return "cauldron_drip_water_pointed_dripstone"
	case 333:
		return "cauldron_drip_lava_pointed_dripstone"
	case 334:
		return "drip_water_pointed_dripstone"
	case 335:
		return "drip_lava_pointed_dripstone"
	case 336:
		return "pick_berries_cave_vines"
	case 337:
		return "tilt_down_big_dripleaf"
	case 338:
		return "tilt_up_big_dripleaf"
	case 339:
		return "copper_wax_on"
	case 340:
		return "copper_wax_off"
	case 341:
		return "scrape"
	case 342:
		return "mob_player_hurt_drown"
	case 343:
		return "mob_player_hurt_on_fire"
	case 344:
		return "mob_player_hurt_freeze"
	case 345:
		return "item_spyglass_use"
	case 346:
		return "item_spyglass_stop_using"
	case 347:
		return "chime_amethyst_block"
	case 348:
		return "ambient_screamer"
	case 349:
		return "hurt_screamer"
	case 350:
		return "death_screamer"
	case 351:
		return "milk_screamer"
	case 352:
		return "jump_to_block"
	case 353:
		return "pre_ram"
	case 354:
		return "pre_ram_screamer"
	case 355:
		return "ram_impact"
	case 356:
		return "ram_impact_screamer"
	case 357:
		return "squid_ink_squirt"
	case 358:
		return "glow_squid_ink_squirt"
	case 359:
		return "convert_to_stray"
	case 360:
		return "cake_add_candle"
	case 361:
		return "extinguish_candle"
	case 362:
		return "ambient_candle"
	case 363:
		return "block_click"
	case 364:
		return "block_click_fail"
	case 365:
		return "block_sculk_catalyst_bloom"
	case 366:
		return "block_sculk_shrieker_shriek"
	case 367:
		return "nearby_close"
	case 368:
		return "nearby_closer"
	case 369:
		return "nearby_closest"
	case 370:
		return "agitated"
	case 371:
		return "record_otherside"
	case 372:
		return "tongue"
	case 373:
		return "irongolem_crack"
	case 374:
		return "irongolem_repair"
	case 375:
		return "listening"
	case 376:
		return "heartbeat"
	case 377:
		return "horn_break"
	case 379:
		return "block_sculk_spread"
	case 380:
		return "charge_sculk"
	case 381:
		return "block_sculk_sensor_place"
	case 382:
		return "block_sculk_shrieker_place"
	case 383:
		return "horn_call0"
	case 384:
		return "horn_call1"
	case 385:
		return "horn_call2"
	case 386:
		return "horn_call3"
	case 387:
		return "horn_call4"
	case 388:
		return "horn_call5"
	case 389:
		return "horn_call6"
	case 390:
		return "horn_call7"
	case 426:
		return "imitate_warden"
	case 427:
		return "listening_angry"
	case 428:
		return "item_given"
	case 429:
		return "item_taken"
	case 430:
		return "disappeared"
	case 431:
		return "reappeared"
	case 432:
		return "drink_milk"
	case 433:
		return "block_frog_spawn_hatch"
	case 434:
		return "lay_spawn"
	case 435:
		return "block_frog_spawn_break"
	case 436:
		return "sonic_boom"
	case 437:
		return "sonic_charge"
	case 438:
		return "item_thrown"
	case 439:
		return "record_5"
	case 440:
		return "convert_to_frog"
	case 442:
		return "block_enchanting_table_use"
	case 443:
		return "step_sand"
	case 444:
		return "dash_ready"
	case 445:
		return "bundle_drop_contents"
	case 446:
		return "bundle_insert"
	case 447:
		return "bundle_remove_one"
	case 448:
		return "pressure_plate_click_off"
	case 449:
		return "pressure_plate_click_on"
	case 450:
		return "button_click_off"
	case 451:
		return "button_click_on"
	case 452:
		return "door_open"
	case 453:
		return "door_close"
	case 454:
		return "trapdoor_open"
	case 455:
		return "trapdoor_close"
	case 456:
		return "fence_gate_open"
	case 457:
		return "fence_gate_close"
	case 458:
		return "insert"
	case 459:
		return "pickup"
	case 460:
		return "insert_enchanted"
	case 461:
		return "pickup_enchanted"
	case 462:
		return "brush"
	case 463:
		return "brush_completed"
	case 464:
		return "shatter_pot"
	case 465:
		return "break_pot"
	case 466:
		return "block_sniffer_egg_crack"
	case 467:
		return "block_sniffer_egg_hatch"
	case 468:
		return "block_sign_waxed_interact_fail"
	case 469:
		return "record_relic"
	case 470:
		return "note_bass"
	case 471:
		return "pumpkin_carve"
	case 472:
		return "mob_husk_convert_to_zombie"
	case 473:
		return "mob_pig_death"
	case 474:
		return "mob_hoglin_converted_to_zombified"
	case 475:
		return "ambient_underwater_enter"
	case 476:
		return "ambient_underwater_exit"
	case 477:
		return "bottle_fill"
	case 478:
		return "bottle_empty"
	case 479:
		return "crafter_craft"
	case 480:
		return "crafter_fail"
	case 481:
		return "block_decorated_pot_insert"
	case 482:
		return "block_decorated_pot_insert_fail"
	case 483:
		return "crafter_disable_slot"
	case 484:
		return "trial_spawner_open_shutter"
	case 485:
		return "trial_spawner_eject_item"
	case 486:
		return "trial_spawner_detect_player"
	case 487:
		return "trial_spawner_spawn_mob"
	case 488:
		return "trial_spawner_close_shutter"
	case 489:
		return "trial_spawner_ambient"
	case 490:
		return "block_copper_bulb_turn_on"
	case 491:
		return "block_copper_bulb_turn_off"
	case 492:
		return "ambient_in_air"
	case 493:
		return "breeze_wind_charge_burst"
	case 494:
		return "imitate_breeze"
	case 495:
		return "mob_armadillo_brush"
	case 496:
		return "mob_armadillo_scute_drop"
	case 497:
		return "armor_equip_wolf"
	case 498:
		return "armor_unequip_wolf"
	case 499:
		return "reflect"
	case 500:
		return "vault_open_shutter"
	case 501:
		return "vault_close_shutter"
	case 502:
		return "vault_eject_item"
	case 503:
		return "vault_insert_item"
	case 504:
		return "vault_insert_item_fail"
	case 505:
		return "vault_ambient"
	case 506:
		return "vault_activate"
	case 507:
		return "vault_deactivate"
	case 508:
		return "hurt_reduced"
	case 509:
		return "wind_charge_burst"
	case 510:
		return "imitate_bogged"
	case 511:
		return "armor_crack_wolf"
	case 512:
		return "armor_break_wolf"
	case 513:
		return "armor_repair_wolf"
	case 514:
		return "mace_smash_air"
	case 515:
		return "mace_smash_ground"
	case 516:
		return "trial_spawner_charge_activate"
	case 517:
		return "trial_spawner_ambient_ominous"
	case 518:
		return "ominous_item_spawner_spawn_item"
	case 519:
		return "ominous_bottle_end_use"
	case 520:
		return "mace_heavy_smash_ground"
	case 521:
		return "ominous_item_spawner_spawn_item_begin"
	case 523:
		return "apply_effect_bad_omen"
	case 524:
		return "apply_effect_raid_omen"
	case 525:
		return "apply_effect_trial_omen"
	case 526:
		return "ominous_item_spawner_about_to_spawn_item"
	case 527:
		return "record_creator"
	case 528:
		return "record_creator_music_box"
	case 529:
		return "record_precipice"
	case 530:
		return "vault_reject_rewarded_player"
	case 531:
		return "imitate_drowned"
	case 532:
		return "imitate_creaking"
	case 533:
		return "bundle_insert_fail"
	case 534:
		return "sponge_absorb"
	case 536:
		return "block_creaking_heart_trail"
	case 537:
		return "creaking_heart_spawn"
	case 538:
		return "activate"
	case 539:
		return "deactivate"
	case 540:
		return "freeze"
	case 541:
		return "unfreeze"
	case 542:
		return "open"
	case 543:
		return "open_long"
	case 544:
		return "close"
	case 545:
		return "close_long"
	case 546:
		return "imitate_phantom"
	case 547:
		return "imitate_zoglin"
	case 548:
		return "imitate_guardian"
	case 549:
		return "imitate_ravager"
	case 550:
		return "imitate_pillager"
	case 551:
		return "place_in_water"
	case 552:
		return "state_change"
	case 553:
		return "imitate_happy_ghast"
	case 554:
		return "armor_unequip_generic"
	case 555:
		return "record_tears"
	case 556:
		return "ambient_weather_the_end_light_flash"
	case 557:
		return "lead_leash"
	case 558:
		return "lead_unleash"
	case 559:
		return "lead_break"
	case 560:
		return "unsaddle"
	case 561:
		return "armor_equip_copper"
	case 562:
		return "record_lava_chicken"
	case 563:
		return "place_item"
	case 564:
		return "single_swap"
	case 565:
		return "multi_swap"
	case 566:
		return "item_enchant_lunge1"
	case 567:
		return "item_enchant_lunge2"
	case 568:
		return "item_enchant_lunge3"
	case 569:
		return "attack_critical"
	case 570:
		return "item_spear_attack_hit"
	case 571:
		return "item_spear_attack_miss"
	case 572:
		return "item_wooden_spear_attack_hit"
	case 573:
		return "item_wooden_spear_attack_miss"
	case 574:
		return "imitate_parched"
	case 575:
		return "imitate_camel_husk"
	case 576:
		return "item_spear_use"
	case 577:
		return "item_wooden_spear_use"
	case 578:
		return "saddle_in_water"
	case 579:
		return "item_stone_spear_attack_hit"
	case 580:
		return "item_iron_spear_attack_hit"
	case 581:
		return "item_copper_spear_attack_hit"
	case 582:
		return "item_golden_spear_attack_hit"
	case 583:
		return "item_diamond_spear_attack_hit"
	case 584:
		return "item_netherite_spear_attack_hit"
	case 585:
		return "item_stone_spear_attack_miss"
	case 586:
		return "item_iron_spear_attack_miss"
	case 587:
		return "item_copper_spear_attack_miss"
	case 588:
		return "item_golden_spear_attack_miss"
	case 589:
		return "item_diamond_spear_attack_miss"
	case 590:
		return "item_netherite_spear_attack_miss"
	case 591:
		return "item_stone_spear_use"
	case 592:
		return "item_iron_spear_use"
	case 593:
		return "item_copper_spear_use"
	case 594:
		return "item_golden_spear_use"
	case 595:
		return "item_diamond_spear_use"
	case 596:
		return "item_netherite_spear_use"
	case 597:
		return "pause_growth"
	case 598:
		return "reset_growth"
	case 599:
		return "pushed_by_player"
	case 600:
		return "bounce"
	default:
		return "unknown"
	}
}

func soundEventFromString(io protocol.IO, x *uint32, s string) {
	switch s {
	case "item_use_on":
		*x = 0
	case "hit":
		*x = 1
	case "step":
		*x = 2
	case "fly":
		*x = 3
	case "jump":
		*x = 4
	case "break":
		*x = 5
	case "place":
		*x = 6
	case "heavy_step":
		*x = 7
	case "gallop":
		*x = 8
	case "fall":
		*x = 9
	case "ambient":
		*x = 10
	case "ambient_baby":
		*x = 11
	case "ambient_in_water":
		*x = 12
	case "breathe":
		*x = 13
	case "death":
		*x = 14
	case "death_in_water":
		*x = 15
	case "death_to_zombie":
		*x = 16
	case "hurt":
		*x = 17
	case "hurt_in_water":
		*x = 18
	case "mad":
		*x = 19
	case "boost":
		*x = 20
	case "bow":
		*x = 21
	case "squish_big":
		*x = 22
	case "squish_small":
		*x = 23
	case "fall_big":
		*x = 24
	case "fall_small":
		*x = 25
	case "splash":
		*x = 26
	case "fizz":
		*x = 27
	case "flap":
		*x = 28
	case "swim":
		*x = 29
	case "drink":
		*x = 30
	case "eat":
		*x = 31
	case "takeoff":
		*x = 32
	case "shake":
		*x = 33
	case "plop":
		*x = 34
	case "land":
		*x = 35
	case "saddle":
		*x = 36
	case "armor":
		*x = 37
	case "mob_armor_stand_place":
		*x = 38
	case "add_chest":
		*x = 39
	case "throw":
		*x = 40
	case "attack":
		*x = 41
	case "attack_nodamage":
		*x = 42
	case "attack_strong":
		*x = 43
	case "warn":
		*x = 44
	case "shear":
		*x = 45
	case "milk":
		*x = 46
	case "thunder":
		*x = 47
	case "explode":
		*x = 48
	case "fire":
		*x = 49
	case "ignite":
		*x = 50
	case "fuse":
		*x = 51
	case "stare":
		*x = 52
	case "spawn":
		*x = 53
	case "shoot":
		*x = 54
	case "break_block":
		*x = 55
	case "launch":
		*x = 56
	case "blast":
		*x = 57
	case "large_blast":
		*x = 58
	case "twinkle":
		*x = 59
	case "remedy":
		*x = 60
	case "unfect":
		*x = 61
	case "levelup":
		*x = 62
	case "bow_hit":
		*x = 63
	case "bullet_hit":
		*x = 64
	case "extinguish_fire":
		*x = 65
	case "item_fizz":
		*x = 66
	case "chest_open":
		*x = 67
	case "chest_closed":
		*x = 68
	case "shulkerbox_open":
		*x = 69
	case "shulkerbox_closed":
		*x = 70
	case "enderchest_open":
		*x = 71
	case "enderchest_closed":
		*x = 72
	case "power_on":
		*x = 73
	case "power_off":
		*x = 74
	case "attach":
		*x = 75
	case "detach":
		*x = 76
	case "deny":
		*x = 77
	case "tripod":
		*x = 78
	case "pop":
		*x = 79
	case "drop_slot":
		*x = 80
	case "note":
		*x = 81
	case "thorns":
		*x = 82
	case "piston_in":
		*x = 83
	case "piston_out":
		*x = 84
	case "portal":
		*x = 85
	case "water":
		*x = 86
	case "lava_pop":
		*x = 87
	case "lava":
		*x = 88
	case "burp":
		*x = 89
	case "bucket_fill_water":
		*x = 90
	case "bucket_fill_lava":
		*x = 91
	case "bucket_empty_water":
		*x = 92
	case "bucket_empty_lava":
		*x = 93
	case "armor_equip_chain":
		*x = 94
	case "armor_equip_diamond":
		*x = 95
	case "armor_equip_generic":
		*x = 96
	case "armor_equip_gold":
		*x = 97
	case "armor_equip_iron":
		*x = 98
	case "armor_equip_leather":
		*x = 99
	case "armor_equip_elytra":
		*x = 100
	case "record_13":
		*x = 101
	case "record_cat":
		*x = 102
	case "record_blocks":
		*x = 103
	case "record_chirp":
		*x = 104
	case "record_far":
		*x = 105
	case "record_mall":
		*x = 106
	case "record_mellohi":
		*x = 107
	case "record_stal":
		*x = 108
	case "record_strad":
		*x = 109
	case "record_ward":
		*x = 110
	case "record_11":
		*x = 111
	case "record_wait":
		*x = 112
	case "stop_record":
		*x = 113
	case "flop":
		*x = 114
	case "elderguardian_curse":
		*x = 115
	case "mob_warning":
		*x = 116
	case "mob_warning_baby":
		*x = 117
	case "teleport":
		*x = 118
	case "shulker_open":
		*x = 119
	case "shulker_close":
		*x = 120
	case "haggle":
		*x = 121
	case "haggle_yes":
		*x = 122
	case "haggle_no":
		*x = 123
	case "haggle_idle":
		*x = 124
	case "chorusgrow":
		*x = 125
	case "chorusdeath":
		*x = 126
	case "glass":
		*x = 127
	case "potion_brewed":
		*x = 128
	case "cast_spell":
		*x = 129
	case "prepare_attack":
		*x = 130
	case "prepare_summon":
		*x = 131
	case "prepare_wololo":
		*x = 132
	case "fang":
		*x = 133
	case "charge":
		*x = 134
	case "camera_take_picture":
		*x = 135
	case "leashknot_place":
		*x = 136
	case "leashknot_break":
		*x = 137
	case "growl":
		*x = 138
	case "whine":
		*x = 139
	case "pant":
		*x = 140
	case "purr":
		*x = 141
	case "purreow":
		*x = 142
	case "death_min_volume":
		*x = 143
	case "death_mid_volume":
		*x = 144
	case "imitate_blaze":
		*x = 145
	case "imitate_cave_spider":
		*x = 146
	case "imitate_creeper":
		*x = 147
	case "imitate_elder_guardian":
		*x = 148
	case "imitate_ender_dragon":
		*x = 149
	case "imitate_enderman":
		*x = 150
	case "imitate_endermite":
		*x = 151
	case "imitate_evocation_illager":
		*x = 152
	case "imitate_ghast":
		*x = 153
	case "imitate_husk":
		*x = 154
	case "imitate_magma_cube":
		*x = 156
	case "imitate_polar_bear":
		*x = 157
	case "imitate_shulker":
		*x = 158
	case "imitate_silverfish":
		*x = 159
	case "imitate_skeleton":
		*x = 160
	case "imitate_slime":
		*x = 161
	case "imitate_spider":
		*x = 162
	case "imitate_stray":
		*x = 163
	case "imitate_vex":
		*x = 164
	case "imitate_vindication_illager":
		*x = 165
	case "imitate_witch":
		*x = 166
	case "imitate_wither":
		*x = 167
	case "imitate_wither_skeleton":
		*x = 168
	case "imitate_wolf":
		*x = 169
	case "imitate_zombie":
		*x = 170
	case "imitate_zombie_pigman":
		*x = 171
	case "imitate_zombie_villager":
		*x = 172
	case "block_end_portal_frame_fill":
		*x = 173
	case "block_end_portal_spawn":
		*x = 174
	case "random_anvil_use":
		*x = 175
	case "bottle_dragonbreath":
		*x = 176
	case "portal_travel":
		*x = 177
	case "item_trident_hit":
		*x = 178
	case "item_trident_return":
		*x = 179
	case "item_trident_riptide_1":
		*x = 180
	case "item_trident_riptide_2":
		*x = 181
	case "item_trident_riptide_3":
		*x = 182
	case "item_trident_throw":
		*x = 183
	case "item_trident_thunder":
		*x = 184
	case "item_trident_hit_ground":
		*x = 185
	case "default":
		*x = 186
	case "block_fletching_table_use":
		*x = 187
	case "elemconstruct_open":
		*x = 188
	case "icebomb_hit":
		*x = 189
	case "balloonpop":
		*x = 190
	case "lt_reaction_icebomb":
		*x = 191
	case "lt_reaction_bleach":
		*x = 192
	case "lt_reaction_epaste":
		*x = 193
	case "lt_reaction_epaste2":
		*x = 194
	case "lt_reaction_glow_stick":
		*x = 195
	case "lt_reaction_glow_stick_2":
		*x = 196
	case "lt_reaction_luminol":
		*x = 197
	case "lt_reaction_salt":
		*x = 198
	case "lt_reaction_fertilizer":
		*x = 199
	case "lt_reaction_fireball":
		*x = 200
	case "lt_reaction_mgsalt":
		*x = 201
	case "lt_reaction_miscfire":
		*x = 202
	case "lt_reaction_fire":
		*x = 203
	case "lt_reaction_miscexplosion":
		*x = 204
	case "lt_reaction_miscmystical":
		*x = 205
	case "lt_reaction_miscmystical2":
		*x = 206
	case "lt_reaction_product":
		*x = 207
	case "sparkler_use":
		*x = 208
	case "glowstick_use":
		*x = 209
	case "sparkler_active":
		*x = 210
	case "convert_to_drowned":
		*x = 211
	case "bucket_fill_fish":
		*x = 212
	case "bucket_empty_fish":
		*x = 213
	case "bubble_up":
		*x = 214
	case "bubble_down":
		*x = 215
	case "bubble_pop":
		*x = 216
	case "bubble_upinside":
		*x = 217
	case "bubble_downinside":
		*x = 218
	case "hurt_baby":
		*x = 219
	case "death_baby":
		*x = 220
	case "step_baby":
		*x = 221
	case "spawn_baby":
		*x = 222
	case "born":
		*x = 223
	case "block_turtle_egg_break":
		*x = 224
	case "block_turtle_egg_crack":
		*x = 225
	case "block_turtle_egg_hatch":
		*x = 226
	case "lay_egg":
		*x = 227
	case "block_turtle_egg_attack":
		*x = 228
	case "beacon_activate":
		*x = 229
	case "beacon_ambient":
		*x = 230
	case "beacon_deactivate":
		*x = 231
	case "beacon_power":
		*x = 232
	case "conduit_activate":
		*x = 233
	case "conduit_ambient":
		*x = 234
	case "conduit_attack":
		*x = 235
	case "conduit_deactivate":
		*x = 236
	case "conduit_short":
		*x = 237
	case "swoop":
		*x = 238
	case "block_bamboo_sapling_place":
		*x = 239
	case "presneeze":
		*x = 240
	case "sneeze":
		*x = 241
	case "ambient_tame":
		*x = 242
	case "scared":
		*x = 243
	case "block_scaffolding_climb":
		*x = 244
	case "crossbow_loading_start":
		*x = 245
	case "crossbow_loading_middle":
		*x = 246
	case "crossbow_loading_end":
		*x = 247
	case "crossbow_shoot":
		*x = 248
	case "crossbow_quick_charge_start":
		*x = 249
	case "crossbow_quick_charge_middle":
		*x = 250
	case "crossbow_quick_charge_end":
		*x = 251
	case "ambient_aggressive":
		*x = 252
	case "ambient_worried":
		*x = 253
	case "cant_breed":
		*x = 254
	case "item_shield_block":
		*x = 255
	case "item_book_put":
		*x = 256
	case "block_grindstone_use":
		*x = 257
	case "block_bell_hit":
		*x = 258
	case "block_campfire_crackle":
		*x = 259
	case "roar":
		*x = 260
	case "stun":
		*x = 261
	case "block_sweet_berry_bush_hurt":
		*x = 262
	case "block_sweet_berry_bush_pick":
		*x = 263
	case "block_cartography_table_use":
		*x = 264
	case "block_stonecutter_use":
		*x = 265
	case "block_composter_empty":
		*x = 266
	case "block_composter_fill":
		*x = 267
	case "block_composter_fill_success":
		*x = 268
	case "block_composter_ready":
		*x = 269
	case "block_barrel_open":
		*x = 270
	case "block_barrel_close":
		*x = 271
	case "raid_horn":
		*x = 272
	case "block_loom_use":
		*x = 273
	case "ambient_in_raid":
		*x = 274
	case "ui_cartography_table_take_result":
		*x = 275
	case "ui_stonecutter_take_result":
		*x = 276
	case "ui_loom_take_result":
		*x = 277
	case "block_smoker_smoke":
		*x = 278
	case "block_blastfurnace_fire_crackle":
		*x = 279
	case "block_smithing_table_use":
		*x = 280
	case "screech":
		*x = 281
	case "sleep":
		*x = 282
	case "block_furnace_lit":
		*x = 283
	case "convert_mooshroom":
		*x = 284
	case "milk_suspiciously":
		*x = 285
	case "celebrate":
		*x = 286
	case "jump_prevent":
		*x = 287
	case "ambient_pollinate":
		*x = 288
	case "block_beehive_drip":
		*x = 289
	case "block_beehive_enter":
		*x = 290
	case "block_beehive_exit":
		*x = 291
	case "block_beehive_work":
		*x = 292
	case "block_beehive_shear":
		*x = 293
	case "drink_honey":
		*x = 294
	case "ambient_cave":
		*x = 295
	case "retreat":
		*x = 296
	case "converted_to_zombified":
		*x = 297
	case "admire":
		*x = 298
	case "step_lava":
		*x = 299
	case "tempt":
		*x = 300
	case "panic":
		*x = 301
	case "angry":
		*x = 302
	case "ambient_warped_forest_mood":
		*x = 303
	case "ambient_soulsand_valley_mood":
		*x = 304
	case "ambient_nether_wastes_mood":
		*x = 305
	case "ambient_basalt_deltas_mood":
		*x = 306
	case "ambient_crimson_forest_mood":
		*x = 307
	case "respawn_anchor_charge":
		*x = 308
	case "respawn_anchor_deplete":
		*x = 309
	case "respawn_anchor_set_spawn":
		*x = 310
	case "respawn_anchor_ambient":
		*x = 311
	case "particle_soul_escape_quiet":
		*x = 312
	case "particle_soul_escape_loud":
		*x = 313
	case "record_pigstep":
		*x = 314
	case "lodestone_compass_link_compass_to_lodestone":
		*x = 315
	case "smithing_table_use":
		*x = 316
	case "armor_equip_netherite":
		*x = 317
	case "ambient_warped_forest_loop":
		*x = 318
	case "ambient_soulsand_valley_loop":
		*x = 319
	case "ambient_nether_wastes_loop":
		*x = 320
	case "ambient_basalt_deltas_loop":
		*x = 321
	case "ambient_crimson_forest_loop":
		*x = 322
	case "ambient_warped_forest_additions":
		*x = 323
	case "ambient_soulsand_valley_additions":
		*x = 324
	case "ambient_nether_wastes_additions":
		*x = 325
	case "ambient_basalt_deltas_additions":
		*x = 326
	case "ambient_crimson_forest_additions":
		*x = 327
	case "power_on_sculk_sensor":
		*x = 328
	case "power_off_sculk_sensor":
		*x = 329
	case "bucket_fill_powder_snow":
		*x = 330
	case "bucket_empty_powder_snow":
		*x = 331
	case "cauldron_drip_water_pointed_dripstone":
		*x = 332
	case "cauldron_drip_lava_pointed_dripstone":
		*x = 333
	case "drip_water_pointed_dripstone":
		*x = 334
	case "drip_lava_pointed_dripstone":
		*x = 335
	case "pick_berries_cave_vines":
		*x = 336
	case "tilt_down_big_dripleaf":
		*x = 337
	case "tilt_up_big_dripleaf":
		*x = 338
	case "copper_wax_on":
		*x = 339
	case "copper_wax_off":
		*x = 340
	case "scrape":
		*x = 341
	case "mob_player_hurt_drown":
		*x = 342
	case "mob_player_hurt_on_fire":
		*x = 343
	case "mob_player_hurt_freeze":
		*x = 344
	case "item_spyglass_use":
		*x = 345
	case "item_spyglass_stop_using":
		*x = 346
	case "chime_amethyst_block":
		*x = 347
	case "ambient_screamer":
		*x = 348
	case "hurt_screamer":
		*x = 349
	case "death_screamer":
		*x = 350
	case "milk_screamer":
		*x = 351
	case "jump_to_block":
		*x = 352
	case "pre_ram":
		*x = 353
	case "pre_ram_screamer":
		*x = 354
	case "ram_impact":
		*x = 355
	case "ram_impact_screamer":
		*x = 356
	case "squid_ink_squirt":
		*x = 357
	case "glow_squid_ink_squirt":
		*x = 358
	case "convert_to_stray":
		*x = 359
	case "cake_add_candle":
		*x = 360
	case "extinguish_candle":
		*x = 361
	case "ambient_candle":
		*x = 362
	case "block_click":
		*x = 363
	case "block_click_fail":
		*x = 364
	case "block_sculk_catalyst_bloom":
		*x = 365
	case "block_sculk_shrieker_shriek":
		*x = 366
	case "nearby_close":
		*x = 367
	case "nearby_closer":
		*x = 368
	case "nearby_closest":
		*x = 369
	case "agitated":
		*x = 370
	case "record_otherside":
		*x = 371
	case "tongue":
		*x = 372
	case "irongolem_crack":
		*x = 373
	case "irongolem_repair":
		*x = 374
	case "listening":
		*x = 375
	case "heartbeat":
		*x = 376
	case "horn_break":
		*x = 377
	case "block_sculk_spread":
		*x = 379
	case "charge_sculk":
		*x = 380
	case "block_sculk_sensor_place":
		*x = 381
	case "block_sculk_shrieker_place":
		*x = 382
	case "horn_call0":
		*x = 383
	case "horn_call1":
		*x = 384
	case "horn_call2":
		*x = 385
	case "horn_call3":
		*x = 386
	case "horn_call4":
		*x = 387
	case "horn_call5":
		*x = 388
	case "horn_call6":
		*x = 389
	case "horn_call7":
		*x = 390
	case "imitate_warden":
		*x = 426
	case "listening_angry":
		*x = 427
	case "item_given":
		*x = 428
	case "item_taken":
		*x = 429
	case "disappeared":
		*x = 430
	case "reappeared":
		*x = 431
	case "drink_milk":
		*x = 432
	case "block_frog_spawn_hatch":
		*x = 433
	case "lay_spawn":
		*x = 434
	case "block_frog_spawn_break":
		*x = 435
	case "sonic_boom":
		*x = 436
	case "sonic_charge":
		*x = 437
	case "item_thrown":
		*x = 438
	case "record_5":
		*x = 439
	case "convert_to_frog":
		*x = 440
	case "block_enchanting_table_use":
		*x = 442
	case "step_sand":
		*x = 443
	case "dash_ready":
		*x = 444
	case "bundle_drop_contents":
		*x = 445
	case "bundle_insert":
		*x = 446
	case "bundle_remove_one":
		*x = 447
	case "pressure_plate_click_off":
		*x = 448
	case "pressure_plate_click_on":
		*x = 449
	case "button_click_off":
		*x = 450
	case "button_click_on":
		*x = 451
	case "door_open":
		*x = 452
	case "door_close":
		*x = 453
	case "trapdoor_open":
		*x = 454
	case "trapdoor_close":
		*x = 455
	case "fence_gate_open":
		*x = 456
	case "fence_gate_close":
		*x = 457
	case "insert":
		*x = 458
	case "pickup":
		*x = 459
	case "insert_enchanted":
		*x = 460
	case "pickup_enchanted":
		*x = 461
	case "brush":
		*x = 462
	case "brush_completed":
		*x = 463
	case "shatter_pot":
		*x = 464
	case "break_pot":
		*x = 465
	case "block_sniffer_egg_crack":
		*x = 466
	case "block_sniffer_egg_hatch":
		*x = 467
	case "block_sign_waxed_interact_fail":
		*x = 468
	case "record_relic":
		*x = 469
	case "note_bass":
		*x = 470
	case "pumpkin_carve":
		*x = 471
	case "mob_husk_convert_to_zombie":
		*x = 472
	case "mob_pig_death":
		*x = 473
	case "mob_hoglin_converted_to_zombified":
		*x = 474
	case "ambient_underwater_enter":
		*x = 475
	case "ambient_underwater_exit":
		*x = 476
	case "bottle_fill":
		*x = 477
	case "bottle_empty":
		*x = 478
	case "crafter_craft":
		*x = 479
	case "crafter_fail":
		*x = 480
	case "block_decorated_pot_insert":
		*x = 481
	case "block_decorated_pot_insert_fail":
		*x = 482
	case "crafter_disable_slot":
		*x = 483
	case "trial_spawner_open_shutter":
		*x = 484
	case "trial_spawner_eject_item":
		*x = 485
	case "trial_spawner_detect_player":
		*x = 486
	case "trial_spawner_spawn_mob":
		*x = 487
	case "trial_spawner_close_shutter":
		*x = 488
	case "trial_spawner_ambient":
		*x = 489
	case "block_copper_bulb_turn_on":
		*x = 490
	case "block_copper_bulb_turn_off":
		*x = 491
	case "ambient_in_air":
		*x = 492
	case "breeze_wind_charge_burst":
		*x = 493
	case "imitate_breeze":
		*x = 494
	case "mob_armadillo_brush":
		*x = 495
	case "mob_armadillo_scute_drop":
		*x = 496
	case "armor_equip_wolf":
		*x = 497
	case "armor_unequip_wolf":
		*x = 498
	case "reflect":
		*x = 499
	case "vault_open_shutter":
		*x = 500
	case "vault_close_shutter":
		*x = 501
	case "vault_eject_item":
		*x = 502
	case "vault_insert_item":
		*x = 503
	case "vault_insert_item_fail":
		*x = 504
	case "vault_ambient":
		*x = 505
	case "vault_activate":
		*x = 506
	case "vault_deactivate":
		*x = 507
	case "hurt_reduced":
		*x = 508
	case "wind_charge_burst":
		*x = 509
	case "imitate_bogged":
		*x = 510
	case "armor_crack_wolf":
		*x = 511
	case "armor_break_wolf":
		*x = 512
	case "armor_repair_wolf":
		*x = 513
	case "mace_smash_air":
		*x = 514
	case "mace_smash_ground":
		*x = 515
	case "trial_spawner_charge_activate":
		*x = 516
	case "trial_spawner_ambient_ominous":
		*x = 517
	case "ominous_item_spawner_spawn_item":
		*x = 518
	case "ominous_bottle_end_use":
		*x = 519
	case "mace_heavy_smash_ground":
		*x = 520
	case "ominous_item_spawner_spawn_item_begin":
		*x = 521
	case "apply_effect_bad_omen":
		*x = 523
	case "apply_effect_raid_omen":
		*x = 524
	case "apply_effect_trial_omen":
		*x = 525
	case "ominous_item_spawner_about_to_spawn_item":
		*x = 526
	case "record_creator":
		*x = 527
	case "record_creator_music_box":
		*x = 528
	case "record_precipice":
		*x = 529
	case "vault_reject_rewarded_player":
		*x = 530
	case "imitate_drowned":
		*x = 531
	case "imitate_creaking":
		*x = 532
	case "bundle_insert_fail":
		*x = 533
	case "sponge_absorb":
		*x = 534
	case "block_creaking_heart_trail":
		*x = 536
	case "creaking_heart_spawn":
		*x = 537
	case "activate":
		*x = 538
	case "deactivate":
		*x = 539
	case "freeze":
		*x = 540
	case "unfreeze":
		*x = 541
	case "open":
		*x = 542
	case "open_long":
		*x = 543
	case "close":
		*x = 544
	case "close_long":
		*x = 545
	case "imitate_phantom":
		*x = 546
	case "imitate_zoglin":
		*x = 547
	case "imitate_guardian":
		*x = 548
	case "imitate_ravager":
		*x = 549
	case "imitate_pillager":
		*x = 550
	case "place_in_water":
		*x = 551
	case "state_change":
		*x = 552
	case "imitate_happy_ghast":
		*x = 553
	case "armor_unequip_generic":
		*x = 554
	case "record_tears":
		*x = 555
	case "ambient_weather_the_end_light_flash":
		*x = 556
	case "lead_leash":
		*x = 557
	case "lead_unleash":
		*x = 558
	case "lead_break":
		*x = 559
	case "unsaddle":
		*x = 560
	case "armor_equip_copper":
		*x = 561
	case "record_lava_chicken":
		*x = 562
	case "place_item":
		*x = 563
	case "single_swap":
		*x = 564
	case "multi_swap":
		*x = 565
	case "item_enchant_lunge1":
		*x = 566
	case "item_enchant_lunge2":
		*x = 567
	case "item_enchant_lunge3":
		*x = 568
	case "attack_critical":
		*x = 569
	case "item_spear_attack_hit":
		*x = 570
	case "item_spear_attack_miss":
		*x = 571
	case "item_wooden_spear_attack_hit":
		*x = 572
	case "item_wooden_spear_attack_miss":
		*x = 573
	case "imitate_parched":
		*x = 574
	case "imitate_camel_husk":
		*x = 575
	case "item_spear_use":
		*x = 576
	case "item_wooden_spear_use":
		*x = 577
	case "saddle_in_water":
		*x = 578
	case "item_stone_spear_attack_hit":
		*x = 579
	case "item_iron_spear_attack_hit":
		*x = 580
	case "item_copper_spear_attack_hit":
		*x = 581
	case "item_golden_spear_attack_hit":
		*x = 582
	case "item_diamond_spear_attack_hit":
		*x = 583
	case "item_netherite_spear_attack_hit":
		*x = 584
	case "item_stone_spear_attack_miss":
		*x = 585
	case "item_iron_spear_attack_miss":
		*x = 586
	case "item_copper_spear_attack_miss":
		*x = 587
	case "item_golden_spear_attack_miss":
		*x = 588
	case "item_diamond_spear_attack_miss":
		*x = 589
	case "item_netherite_spear_attack_miss":
		*x = 590
	case "item_stone_spear_use":
		*x = 591
	case "item_iron_spear_use":
		*x = 592
	case "item_copper_spear_use":
		*x = 593
	case "item_golden_spear_use":
		*x = 594
	case "item_diamond_spear_use":
		*x = 595
	case "item_netherite_spear_use":
		*x = 596
	case "pause_growth":
		*x = 597
	case "reset_growth":
		*x = 598
	case "pushed_by_player":
		*x = 599
	case "bounce":
		*x = 600
	default:
		io.InvalidValue(s, "soundType", "unknown sound")
	}
}
