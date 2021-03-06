package ogame

type rocketLauncher struct {
	BaseDefense
}

func newRocketLauncher() *rocketLauncher {
	d := new(rocketLauncher)
	d.Name = "rocket launcher"
	d.ID = RocketLauncherID
	d.Price = Resources{Metal: 2000}
	d.StructuralIntegrity = 2000
	d.ShieldPower = 20
	d.WeaponPower = 80
	d.RapidfireFrom = map[ID]int{BomberID: 20, CruiserID: 10, DeathstarID: 200}
	d.Requirements = map[ID]int{ShipyardID: 1}
	return d
}
