package x00_simple_factory

import "testing"

func TestMethodFactory(t *testing.T) {
	factoryA := &FactoryA{}
	factoryB := &FactoryB{}

	factoryA.createFactory().Use()
	factoryB.createFactory().Use()
}
