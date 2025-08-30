package weatherapi

import "fullcycle-lab-3/internal/app"

type GetCurretModel struct {
	Location LocationModel
	Current  CurrentModel
}

type LocationModel struct {
	Name    string
	Region  string
	Country string
}

type CurrentModel struct {
	Temp_C   float32
	Temp_F   float32
	Wind_Kph float32
}

func (m *GetCurretModel) Map() *app.ClimaServiceModel {
	return &app.ClimaServiceModel{
		Temp_C: m.Current.Temp_C,
	}
}
