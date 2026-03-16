package geometry_calculator

type GeometryFigure interface {
	CalculateArea() int
	CalculatePerimeter() int
}

type GeometryModule struct {
	geometryFigure GeometryFigure
}

func NewGeometryModule(geometryFigure GeometryFigure) *GeometryModule {
	return &GeometryModule{
		geometryFigure: geometryFigure,
	}
}

func (g *GeometryModule) CalculateArea() int {
	return g.geometryFigure.CalculateArea()
}

func (g *GeometryModule) CalculatePerimeter() int {
	return g.geometryFigure.CalculatePerimeter()
}
