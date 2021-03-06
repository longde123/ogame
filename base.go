package ogame

// Base struct for all ogame objects
type Base struct {
	ID           ID
	Name         string
	Requirements map[ID]int
}

// GetID returns the ogame id of the object
func (b Base) GetID() ID {
	return b.ID
}

// GetName returns the printable name of the object
func (b Base) GetName() string {
	return b.Name
}

// GetRequirements returns the requirements to have this object available
func (b Base) GetRequirements() map[ID]int {
	return b.Requirements
}

// IsAvailable returns either or not the object is available to us
func (b Base) IsAvailable(t CelestialType, resourcesBuildings ResourcesBuildings, facilities Facilities, researches Researches, energy int) bool {
	if t != PlanetType && t != MoonType {
		return false
	}
	if t == PlanetType {
		if b.ID == LunarBaseID ||
			b.ID == SensorPhalanxID ||
			b.ID == JumpGateID {
			return false
		}
	} else if t == MoonType {
		if b.ID == MetalMineID ||
			b.ID == CrystalMineID ||
			b.ID == DeuteriumSynthesizerID ||
			b.ID == SolarPlantID ||
			b.ID == FusionReactorID ||
			b.ID == ResearchLabID ||
			b.ID == AllianceDepotID ||
			b.ID == MissileSiloID ||
			b.ID == NaniteFactoryID ||
			b.ID == TerraformerID ||
			b.ID == SpaceDockID {
			return false
		}
	}
	if b.ID == GravitonTechnologyID && energy < 300000 {
		return false
	}
	type requirement struct {
		ID  ID
		Lvl int
	}
	q := make([]requirement, 0)
	for id, levelNeeded := range b.Requirements {
		q = append(q, requirement{id, levelNeeded})
	}
	for len(q) > 0 {
		var req requirement
		req, q = q[0], q[1:]
		reqs := Objs.ByID(req.ID).GetRequirements()
		for k, v := range reqs {
			q = append(q, requirement{k, v})
		}
		id := req.ID
		levelNeeded := req.Lvl
		if id.IsResourceBuilding() {
			if resourcesBuildings.ByID(id) < levelNeeded {
				return false
			}
		} else if id.IsFacility() {
			if facilities.ByID(id) < levelNeeded {
				return false
			}
		} else if id.IsTech() {
			if researches.ByID(id) < levelNeeded {
				return false
			}
		}
	}
	return true
}
