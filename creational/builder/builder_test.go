package builder

import "testing"

func TestManufacturer(t *testing.T)  {
	manufacturer := Manufacturer{}

	homePCBuilder := &HomeEditionPCBuilder{}
	manufacturer.SetBuilder(homePCBuilder)
	manufacturer.ConstructPC()
	homePC := homePCBuilder.GetPC()

	if homePC.cpu != "i3" {
		t.Errorf("Home edition PC cpu shoul be 'i3' but found %s", homePC.cpu)
	}

	if homePC.gpu != "Intel Graphic Card" {
		t.Errorf("Home edition PC gpu shoul be 'Intel Graphic Card' but found %s", homePC.cpu)
	}

	gamePCBuilder := &GameEditionPCBuilder{}
	manufacturer.SetBuilder(gamePCBuilder)
	manufacturer.ConstructPC()
	gamePC := gamePCBuilder.GetPC()

	if gamePC.cpu != "i7" {
		t.Errorf("Game edition PC cpu shoul be 'i7' but found %s", gamePC.cpu)
	}

	if gamePC.gpu != "AMD Radeon X80" {
		t.Errorf("Game edition PC gpu shoul be 'Intel Graphic Card' but found %s", gamePC.cpu)
	}
}