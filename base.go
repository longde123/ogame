package ogame

import (
	"math"
)

// BaseShip ...
type BaseShip struct {
	ID                  ID
	StructuralIntegrity int
	ShieldPower         int
	WeaponPower         int
	CargoCapacity       int
	BaseSpeed           int
	FuelConsumption     int
	RapidfireFrom       map[ID]int
	RapidfireAgainst    map[ID]int
	Requirements        map[ID]int
	Price               Resources
}

// GetOGameID ...
func (b BaseShip) GetOGameID() ID {
	return b.ID
}

// GetStructuralIntegrity ...
func (b BaseShip) GetStructuralIntegrity() int {
	return b.StructuralIntegrity
}

// GetShieldPower ...
func (b BaseShip) GetShieldPower() int {
	return b.ShieldPower
}

// GetWeaponPower ...
func (b BaseShip) GetWeaponPower() int {
	return b.WeaponPower
}

// GetCargoCapacity ...
func (b BaseShip) GetCargoCapacity() int {
	return b.CargoCapacity
}

// GetBaseSpeed ...
func (b BaseShip) GetBaseSpeed() int {
	return b.BaseSpeed
}

// GetSpeed ...
func (b BaseShip) GetSpeed(techs Researches) int {
	techDriveLvl := 0
	if minLvl, ok := b.Requirements[CombustionDrive.ID]; ok {
		techDriveLvl = techs.CombustionDrive
		if techDriveLvl < minLvl {
			techDriveLvl = minLvl
		}
	} else if minLvl, ok := b.Requirements[ImpulseDrive.ID]; ok {
		techDriveLvl = techs.ImpulseDrive
		if techDriveLvl < minLvl {
			techDriveLvl = minLvl
		}
	} else if minLvl, ok := b.Requirements[HyperspaceDrive.ID]; ok {
		techDriveLvl = techs.HyperspaceDrive
		if techDriveLvl < minLvl {
			techDriveLvl = minLvl
		}
	}
	return int(float64(b.BaseSpeed) + (float64(b.BaseSpeed)*0.2)*float64(techDriveLvl))
}

// GetFuelConsumption ...
func (b BaseShip) GetFuelConsumption() int {
	return b.FuelConsumption
}

// GetRapidfireFrom ...
func (b BaseShip) GetRapidfireFrom() map[ID]int {
	return b.RapidfireFrom
}

// GetRapidfireAgainst ...
func (b BaseShip) GetRapidfireAgainst() map[ID]int {
	return b.RapidfireAgainst
}

// GetPrice ...
func (b BaseShip) GetPrice(int) Resources {
	return b.Price
}

// ConstructionTime ...
func (b BaseShip) ConstructionTime(nbr, universeSpeed int, facilities Facilities) int {
	shipyardLvl := float64(facilities.Shipyard)
	naniteLvl := float64(facilities.NaniteFactory)
	hours := float64(b.StructuralIntegrity) / (2500 * (1 + shipyardLvl) * float64(universeSpeed) * math.Pow(2, naniteLvl))
	secs := hours * 3600
	return int(math.Floor(secs)) * nbr
}

// GetRequirements ...
func (b BaseShip) GetRequirements() map[ID]int {
	return b.Requirements
}

// IsAvailable ...
func (b BaseShip) IsAvailable(_ ResourcesBuildings, facilities Facilities, researches Researches, _ int) bool {
	for ogameID, levelNeeded := range b.Requirements {
		if ogameID.IsFacility() {
			if facilities.ByOGameID(ogameID) < levelNeeded {
				return false
			}
		} else if ogameID.IsTech() {
			if researches.ByOGameID(ogameID) < levelNeeded {
				return false
			}
		}
	}
	return true
}