package proto

import (
	"fmt"
	"image/color"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalPrimitiveShape(io protocol.IO, x *protocol.PrimitiveShape) {
	if IsProtoLT(io, ID860) {
		marshalFlatPrimitiveShape(io, x)
		return
	}
	io.Varuint64(&x.NetworkID)
	protocol.OptionalFunc(io, &x.Type, io.Uint8)
	protocol.OptionalFunc(io, &x.Location, io.Vec3)
	protocol.OptionalFunc(io, &x.Scale, io.Float32)
	protocol.OptionalFunc(io, &x.Rotation, io.Vec3)
	protocol.OptionalFunc(io, &x.TotalTimeLeft, io.Float32)
	if IsProtoGTE(io, ID975) {
		protocol.OptionalFunc(io, &x.MaxRenderDistance, io.Float32)
	}
	protocol.OptionalFunc(io, &x.Colour, io.BEARGB)
	if IsProtoLT(io, ID924) {
		dimensionID, _ := x.DimensionID.Value()
		io.Varint32(&dimensionID)
		x.DimensionID = protocol.Option(dimensionID)
	} else {
		protocol.OptionalFunc(io, &x.DimensionID, io.Varint32)
	}
	if IsProtoGTE(io, ID2168) {
		protocol.OptionalFunc(io, &x.AttachedToEntityID, io.ActorUniqueID)
	} else if IsProtoGTE(io, ID975) {
		present := false
		value, ok := x.AttachedToEntityID.Value()
		present = ok
		io.Bool(&present)
		if present {
			io.Varint64(&value)
			x.AttachedToEntityID = protocol.Option(value)
		}
	} else if IsProtoGTE(io, ID924) {
		value, present := x.AttachedToEntityID.Value()
		io.Bool(&present)
		if present {
			unsigned := uint64(value)
			io.Varuint64(&unsigned)
			x.AttachedToEntityID = protocol.Option(int64(unsigned))
		}
	}
	MarshalShapeData(io, &x.ExtraShapeData)
}

func marshalFlatPrimitiveShape(io protocol.IO, x *protocol.PrimitiveShape) {
	if IsProto(io, ID818) {
		networkID := uint32(x.NetworkID)
		io.Varuint32(&networkID)
		x.NetworkID = uint64(networkID)
	} else {
		io.Varuint64(&x.NetworkID)
	}
	protocol.OptionalFunc(io, &x.Type, io.Uint8)
	protocol.OptionalFunc(io, &x.Location, io.Vec3)
	protocol.OptionalFunc(io, &x.Scale, io.Float32)
	protocol.OptionalFunc(io, &x.Rotation, io.Vec3)
	protocol.OptionalFunc(io, &x.TotalTimeLeft, io.Float32)
	if IsProto(io, ID818) {
		colour, present := x.Colour.Value()
		io.Bool(&present)
		if present {
			argb := int32(colour.A)<<24 | int32(colour.R)<<16 | int32(colour.G)<<8 | int32(colour.B)
			io.Int32(&argb)
			x.Colour = protocol.Option(color.RGBA{A: byte(argb >> 24), R: byte(argb >> 16), G: byte(argb >> 8), B: byte(argb)})
		}
	} else {
		protocol.OptionalFunc(io, &x.Colour, io.BEARGB)
	}
	var text protocol.Optional[string]
	var box, line protocol.Optional[mgl32.Vec3]
	var arrowLength, arrowRadius protocol.Optional[float32]
	var segments protocol.Optional[byte]
	if !IsReader(io) {
		switch shape := x.ExtraShapeData.(type) {
		case *protocol.TextShape:
			text = protocol.Option(shape.Text)
		case *protocol.BoxShape:
			box = protocol.Option(shape.BoxBound)
		case *protocol.LineShape:
			line = protocol.Option(shape.LineEndLocation)
		case *protocol.ArrowShape:
			arrowLength = shape.ArrowHeadLength
			arrowRadius = shape.ArrowHeadRadius
			segments = shape.Segments
		case *protocol.SphereShape:
			segments = protocol.Option(shape.Segments)
		}
	}
	protocol.OptionalFunc(io, &text, io.String)
	protocol.OptionalFunc(io, &box, io.Vec3)
	protocol.OptionalFunc(io, &line, io.Vec3)
	protocol.OptionalFunc(io, &arrowLength, io.Float32)
	protocol.OptionalFunc(io, &arrowRadius, io.Float32)
	if IsProto(io, ID818) {
		value, present := segments.Value()
		io.Bool(&present)
		if present {
			legacySegments := int32(value)
			io.Int32(&legacySegments)
			segments = protocol.Option(byte(legacySegments))
		}
	} else {
		protocol.OptionalFunc(io, &segments, io.Uint8)
	}
	if IsReader(io) {
		switch {
		case optionalPresent(text):
			value, _ := text.Value()
			x.ExtraShapeData = &protocol.TextShape{Text: value}
		case optionalPresent(box):
			value, _ := box.Value()
			x.ExtraShapeData = &protocol.BoxShape{BoxBound: value}
		case optionalPresent(line):
			value, _ := line.Value()
			x.ExtraShapeData = &protocol.LineShape{LineEndLocation: value}
		default:
			x.ExtraShapeData = &protocol.LastShape{}
		}
	}
}

func optionalPresent[T any](value protocol.Optional[T]) bool {
	_, ok := value.Value()
	return ok
}

func MarshalShapeData(io protocol.IO, x *protocol.ShapeData) {
	var shapeType uint32
	if IsReader(io) {
		io.Varuint32(&shapeType)
		switch shapeType {
		case protocol.ShapeDataLast:
			*x = &protocol.LastShape{}
		case protocol.ShapeDataArrow:
			*x = &protocol.ArrowShape{}
		case protocol.ShapeDataText:
			*x = &protocol.TextShape{}
		case protocol.ShapeDataBox:
			*x = &protocol.BoxShape{}
		case protocol.ShapeDataLine:
			*x = &protocol.LineShape{}
		case protocol.ShapeDataSphere:
			*x = &protocol.SphereShape{}
		case protocol.ShapeDataCylinder:
			*x = &protocol.CylinderShape{}
		case protocol.ShapeDataPyramid:
			*x = &protocol.PyramidShape{}
		case protocol.ShapeDataEllipsoid:
			*x = &protocol.EllipsoidShape{}
		case protocol.ShapeDataCone:
			*x = &protocol.ConeShape{}
		default:
			io.UnknownEnumOption(shapeType, "debug shape data type")
			return
		}
	} else {
		switch (*x).(type) {
		case *protocol.LastShape:
			shapeType = protocol.ShapeDataLast
		case *protocol.ArrowShape:
			shapeType = protocol.ShapeDataArrow
		case *protocol.TextShape:
			shapeType = protocol.ShapeDataText
		case *protocol.BoxShape:
			shapeType = protocol.ShapeDataBox
		case *protocol.LineShape:
			shapeType = protocol.ShapeDataLine
		case *protocol.SphereShape:
			shapeType = protocol.ShapeDataSphere
		case *protocol.CylinderShape:
			shapeType = protocol.ShapeDataCylinder
		case *protocol.PyramidShape:
			shapeType = protocol.ShapeDataPyramid
		case *protocol.EllipsoidShape:
			shapeType = protocol.ShapeDataEllipsoid
		case *protocol.ConeShape:
			shapeType = protocol.ShapeDataCone
		default:
			io.UnknownEnumOption(fmt.Sprintf("%T", *x), "debug shape data type")
			return
		}
		io.Varuint32(&shapeType)
	}
	marshalShapePayload(io, *x)
}

func marshalShapePayload(io protocol.IO, x protocol.ShapeData) {
	switch shape := x.(type) {
	case *protocol.LastShape:
	case *protocol.LineShape:
		io.Vec3(&shape.LineEndLocation)
	case *protocol.TextShape:
		io.String(&shape.Text)
		if IsProtoGTE(io, ID975) {
			io.Bool(&shape.UseRotation)
			protocol.OptionalFunc(io, &shape.BackgroundColour, io.BEARGB)
			io.Bool(&shape.DepthTest)
			io.Bool(&shape.ShowBackface)
			io.Bool(&shape.ShowBackfaceText)
		}
	case *protocol.BoxShape:
		io.Vec3(&shape.BoxBound)
	case *protocol.SphereShape:
		io.Uint8(&shape.Segments)
	case *protocol.ArrowShape:
		protocol.OptionalFunc(io, &shape.ArrowEndLocation, io.Vec3)
		protocol.OptionalFunc(io, &shape.ArrowHeadLength, io.Float32)
		protocol.OptionalFunc(io, &shape.ArrowHeadRadius, io.Float32)
		protocol.OptionalFunc(io, &shape.Segments, io.Uint8)
	case *protocol.CylinderShape:
		io.Vec2(&shape.RadiusX)
		io.Vec2(&shape.RadiusZ)
		io.Float32(&shape.Height)
		io.Uint8(&shape.NumSegments)
	case *protocol.PyramidShape:
		io.Float32(&shape.Width)
		protocol.OptionalFunc(io, &shape.Depth, io.Float32)
		io.Float32(&shape.Height)
	case *protocol.EllipsoidShape:
		io.Vec3(&shape.Radii)
		io.Uint8(&shape.SegmentsPerAxis)
	case *protocol.ConeShape:
		io.Vec2(&shape.Radii)
		io.Float32(&shape.Height)
		io.Uint8(&shape.NumSegments)
	}
}
