package postal

import (
	"fmt"
	"strconv"
)

type (
	State     string
	StateName string
)

const (
	Alaska     State     = "AK"
	AlaskaName StateName = "Alaska"

	Alabama     State     = "AL"
	AlabamaName StateName = "Alabama"

	Arkansas     State     = "AR"
	ArkansasName StateName = "Arkansas"

	Arizona     State     = "AZ"
	ArizonaName StateName = "Arizona"

	California     State     = "CA"
	CaliforniaName StateName = "California"

	Colorado     State     = "CO"
	ColoradoName StateName = "Colorado"

	Connecticut     State     = "CT"
	ConnecticutName StateName = "Connecticut"

	WashingtonDC     State     = "DC"
	WashingtonDCName StateName = "Washington DC"

	Delaware     State     = "DE"
	DelawareName StateName = "Delaware"

	Florida     State     = "FL"
	FloridaName StateName = "Florida"

	Georgia     State     = "GA"
	GeorgiaName StateName = "Georgia"

	Hawaii     State     = "HI"
	HawaiiName StateName = "Hawaii"

	Iowa     State     = "IA"
	IowaName StateName = "Iowa"

	Idaho     State     = "ID"
	IdahoName StateName = "Idaho"

	Illinois     State     = "IL"
	IllinoisName StateName = "Illinois"

	Indiana     State     = "IN"
	IndianaName StateName = "Indiana"

	Kansas     State     = "KS"
	KansasName StateName = "Kansas"

	Kentucky     State     = "KY"
	KentuckyName StateName = "Kentucky"

	Louisiana     State     = "LA"
	LouisianaName StateName = "Louisiana"

	Massachusetts     State     = "MA"
	MassachusettsName StateName = "Massachusetts"

	Maryland     State     = "MD"
	MarylandName StateName = "Maryland"

	Maine     State     = "ME"
	MaineName StateName = "Maine"

	Michigan     State     = "MI"
	MichiganName StateName = "Michigan"

	Minnesota     State     = "MN"
	MinnesotaName StateName = "Minnesota"

	Missouri     State     = "MO"
	MissouriName StateName = "Missouri"

	Mississippi     State     = "MS"
	MississippiName StateName = "Mississippi"

	Montana     State     = "MT"
	MontanaName StateName = "Montana"

	NorthCarolina     State     = "NC"
	NorthCarolinaName StateName = "North Carolina"

	NorthDakota     State     = "ND"
	NorthDakotaName StateName = "North Dakota"

	Nebraska     State     = "NE"
	NebraskaName StateName = "Nebraska"

	NewHampshire     State     = "NH"
	NewHampshireName StateName = "New Hampshire"

	NewJersey     State     = "NJ"
	NewJerseyName StateName = "New Jersey"

	NewMexico     State     = "NM"
	NewMexicoName StateName = "New Mexico"

	Nevada     State     = "NV"
	NevadaName StateName = "Nevada"

	NewYork     State     = "NY"
	NewYorkName StateName = "New York"

	Ohio     State     = "OH"
	OhioName StateName = "Ohio"

	Oklahoma     State     = "OK"
	OklahomaName StateName = "Oklahoma"

	Oregon     State     = "OR"
	OregonName StateName = "Oregon"

	Pennsylvania     State     = "PA"
	PennsylvaniaName StateName = "Pennsylvania"

	RhodeIsland     State     = "RI"
	RhodeIslandName StateName = "Rhode Island"

	SouthCarolina     State     = "SC"
	SouthCarolinaName StateName = "South Carolina"

	SouthDakota     State     = "SD"
	SouthDakotaName StateName = "South Dakota"

	Tennessee     State     = "TN"
	TennesseeName StateName = "Tennessee"

	Texas     State     = "TX"
	TexasName StateName = "Texas"

	Utah     State     = "UT"
	UtahName StateName = "Utah"

	Virginia     State     = "VA"
	VirginiaName StateName = "Virginia"

	Vermont     State     = "VT"
	VermontName StateName = "Vermont"

	Washington     State     = "WA"
	WashingtonName StateName = "Washington"

	Wisconsin     State     = "WI"
	WisconsinName StateName = "Wisconsin"

	WestVirginia     State     = "WV"
	WestVirginiaName StateName = "West Virginia"

	Wyoming     State     = "WY"
	WyomingName StateName = "Wyoming"
)

func (s State) Is(state State) bool {
	return s == state
}

func (s State) Name() StateName {
	switch s {
	case Alaska:
		return AlaskaName
	case Alabama:
		return AlabamaName
	case Arkansas:
		return ArkansasName
	case Arizona:
		return ArizonaName
	case California:
		return CaliforniaName
	case Colorado:
		return ColoradoName
	case Connecticut:
		return ConnecticutName
	case WashingtonDC:
		return WashingtonDCName
	case Delaware:
		return DelawareName
	case Florida:
		return FloridaName
	case Georgia:
		return GeorgiaName
	case Hawaii:
		return HawaiiName
	case Iowa:
		return IowaName
	case Idaho:
		return IdahoName
	case Illinois:
		return IllinoisName
	case Indiana:
		return IndianaName
	case Kansas:
		return KansasName
	case Kentucky:
		return KentuckyName
	case Louisiana:
		return LouisianaName
	case Massachusetts:
		return MassachusettsName
	case Maryland:
		return MarylandName
	case Maine:
		return MaineName
	case Michigan:
		return MichiganName
	case Minnesota:
		return MinnesotaName
	case Missouri:
		return MissouriName
	case Mississippi:
		return MississippiName
	case Montana:
		return MontanaName
	case NorthCarolina:
		return NorthCarolinaName
	case NorthDakota:
		return NorthDakotaName
	case Nebraska:
		return NebraskaName
	case NewHampshire:
		return NewHampshireName
	case NewJersey:
		return NewJerseyName
	case NewMexico:
		return NewMexicoName
	case Nevada:
		return NevadaName
	case NewYork:
		return NewYorkName
	case Ohio:
		return OhioName
	case Oklahoma:
		return OklahomaName
	case Oregon:
		return OregonName
	case Pennsylvania:
		return PennsylvaniaName
	case RhodeIsland:
		return RhodeIslandName
	case SouthCarolina:
		return SouthCarolinaName
	case SouthDakota:
		return SouthDakotaName
	case Tennessee:
		return TennesseeName
	case Texas:
		return TexasName
	case Utah:
		return UtahName
	case Virginia:
		return VirginiaName
	case Vermont:
		return VermontName
	case Washington:
		return WashingtonName
	case Wisconsin:
		return WisconsinName
	case WestVirginia:
		return WestVirginiaName
	case Wyoming:
		return WyomingName
	}

	return ""
}

func FindState(postal string) (State, error) {
	if !IsValid(postal) {
		return "", fmt.Errorf("invalid code '%s': %w", postal, ErrInvalidCode)
	}

	code, err := strconv.ParseInt(postal, 10, 64)
	if err != nil {
		return "", fmt.Errorf("code is not a valid number: %w", ErrInvalidCode)
	}

	if code < 0 {
		return "", ErrStateNotFound
	}

	state := getStateFromCode(code)
	if state == "" {
		return "", ErrStateNotFound
	}

	return state, nil
}

func getStateFromCode(code int64) State {
	switch {
	case rng(code, 99500, 99999):
		return Alaska
	case rng(code, 35000, 36999):
		return Alabama
	case rng(code, 71600, 72999):
		return Arkansas
	case rng(code, 85000, 86999):
		return Arizona
	case rng(code, 90000, 96699):
		return California
	case rng(code, 80000, 81999):
		return Colorado
	case rng(code, 6000, 6999):
		return Connecticut
	case rng(code, 20000, 20599):
		return Washington
	case rng(code, 19700, 19999):
		return Delaware
	case rng(code, 32000, 34999):
		return Florida
	case rng(code, 30000, 31999):
		return Georgia
	case rng(code, 96700, 96999):
		return Hawaii
	case rng(code, 50000, 52999):
		return Iowa
	case rng(code, 83200, 83999):
		return Idaho
	case rng(code, 60000, 62999):
		return Illinois
	case rng(code, 46000, 47999):
		return Indiana
	case rng(code, 66000, 67999):
		return Kansas
	case rng(code, 40000, 42999):
		return Kentucky
	case rng(code, 70000, 71599):
		return Louisiana
	case rng(code, 1000, 2799):
		return Massachusetts
	case rng(code, 20600, 21999):
		return Maryland
	case rng(code, 3900, 4999):
		return Maine
	case rng(code, 48000, 49999):
		return Michigan
	case rng(code, 55000, 56999):
		return Minnesota
	case rng(code, 63000, 65999):
		return Missouri
	case rng(code, 38600, 39999):
		return Mississippi
	case rng(code, 59000, 59999):
		return Montana
	case rng(code, 27000, 28999):
		return NorthCarolina
	case rng(code, 58000, 58999):
		return NorthDakota
	case rng(code, 68000, 69999):
		return Nebraska
	case rng(code, 3000, 3899):
		return NewHampshire
	case rng(code, 7000, 8999):
		return NewJersey
	case rng(code, 87000, 88499):
		return NewMexico
	case rng(code, 88900, 89999):
		return Nevada
	case rng(code, 10000, 14999):
		return NewYork
	case rng(code, 43000, 45999):
		return Ohio
	case rng(code, 73000, 74999):
		return Oklahoma
	case rng(code, 97000, 97999):
		return Oregon
	case rng(code, 15000, 19699):
		return Pennsylvania
	case rng(code, 2800, 2999):
		return RhodeIsland
	case rng(code, 29000, 29999):
		return SouthCarolina
	case rng(code, 57000, 57999):
		return SouthDakota
	case rng(code, 37000, 38599):
		return Tennessee
	case rng(code, 75000, 79999) || rng(code, 88500, 88599):
		return Texas
	case rng(code, 84000, 84999):
		return Utah
	case rng(code, 22000, 24699):
		return Virginia
	case rng(code, 5000, 5999):
		return Vermont
	case rng(code, 20000, 20599):
		return Washington
	case rng(code, 98000, 99499):
		return Wisconsin
	case rng(code, 24700, 26999):
		return WestVirginia
	case rng(code, 82000, 83199):
		return Wyoming
	}

	return ""
}

func rng(target, min, max int64) bool {
	return target >= min && target <= max
}
