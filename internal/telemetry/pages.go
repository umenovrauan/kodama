package telemetry

type pageStatic struct {
	SmVersion           [15]uint16
	AcVersion           [15]uint16
	NumberOfSessions    int32
	NumCars             int32
	CarModel            [33]uint16
	Track               [33]uint16
	PlayerName          [33]uint16
	PlayerSurname       [33]uint16
	PlayerNick          [33]uint16
	SectorCount         int32
	_                   float32 // Unused maxTorque
	_                   float32 // Unused maxPower
	MaxRpm              int32
	MaxFuel             float32
	_                   [4]float32 // Unused suspensionMaxTravel
	_                   [4]float32 // Unused tyreRadius
	_                   float32    // Unused maxTurboBoost
	_                   float32    // Unused deprecated_1
	_                   float32    // Unused deprecated_2
	PenaltiesEnabled    int32
	AidFuelRate         float32
	AidTireRate         float32
	AidMechanicalDamage float32
	AllowTyreBlankets   int32
	AidStability        float32
	AidAutoClutch       int32
	AidAutoBlip         int32
	_                   int32      // Unused hasDRS
	_                   int32      // Unused hasERS
	_                   int32      // Unused hasKERS
	_                   float32    // Unused kersMaxJ
	_                   int32      // Unused engineBrakeSettingCount
	_                   int32      // Unused ersPowerControllerCount
	_                   float32    // Unused traclSpLineLenght
	_                   [33]uint16 // Unused trackConfiguration
	_                   float32    // Unused ersMaxJ
	_                   int32      // Unused isTimedRace
	_                   int32      // Unused hasExtraLap
	_                   [33]uint16 // Unused carSkin
	_                   int32      // Unused reverseGridPositions
	PitWindowStart      int32
	PitWindowEnd        int32
	IsOnline            int32
	DryTyresName        [33]uint16
	WetTyresName        [33]uint16
}

type pagePhysics struct {
	PacketID             int32
	Gas                  float32
	Brake                float32
	Fuel                 float32
	Gear                 int32
	RPM                  int32
	SteerAngle           float32
	SpeedKmh             float32
	Velocity             [3]float32
	AccG                 [3]float32
	WheelSlip            [4]float32
	_                    [4]float32 // Unused wheelload
	WheelPressure        [4]float32
	WheelAngularSpeed    [4]float32
	_                    [4]float32 // Unused tyre wear
	_                    [4]float32 // unused tyredirtylevel
	TireCoreTemperature  [4]float32
	_                    [4]float32 // Unused camberRAD
	SuspensionTravel     [4]float32
	_                    float32 // Unused drs
	TractionControl      float32
	Yaw                  float32
	Pitch                float32
	Roll                 float32
	_                    float32 // Unused cgHeight
	CarDamage            [5]float32
	_                    int32 // Unused number of tyres out
	PitLimiterOn         int32
	Abs                  float32
	_                    float32 // Unused Kers charge
	_                    float32 // Unused Kers Input
	AutoShifterOn        int32
	_                    [2]float32 // Unused ride height
	TurboBoost           float32
	_                    float32 // Unused ballast
	_                    float32 // Unused air density
	AirTemperature       float32
	RoadTemperature      float32
	LocalAngularVelocity [3]float32
	ForceFeedback        float32
	_                    float32 // Unused performanceMeter
	_                    float32 // Unused engineBrake
	_                    float32 // Unused ersRecoveryLevel
	_                    float32 // Unused ersPowerLevel
	_                    int32   // Unused ersHeatCharging
	_                    int32   // Unused ersIsCharging
	_                    float32 // Unused kersCurrentKJ
	_                    int32   // Unused drsAvailable
	_                    int32   // Unused drsEnabled
	BrakeTemperature     [4]float32
	Clutch               float32
	_                    [4]float32 // Unused tyre tempI
	_                    [4]float32 // Unused tyre tempM
	_                    [4]float32 // Unused tyre tempO
	IsAIControlled       int32
	TyreContactPoint     [4][3]float32
	TyreContactNormal    [4][3]float32
	TyreContactHeading   [4][3]float32
	BrakeBias            float32
	LocalVelocity        [3]float32
	_                    int32      // Unused P2P activated
	_                    int32      // Unused P2P status
	_                    int32      // Unused current maxrpm
	_                    [4]float32 // Unused mz
	_                    [4]float32 // Unused fx
	_                    [4]float32 // Unused fy
	SlipRatio            [4]float32
	SlipAngle            [4]float32
	_                    int32      // Unused tcinAction
	_                    int32      // Unused AbsinAction
	_                    [4]float32 // Unused suspensionDamage
	_                    [4]float32 // Unused tyretemp
	WaterTemp            float32
	BrakePressure        [4]float32
	FrontBrakeCompound   int32
	RearBrakeCompound    int32
	PadLife              [4]float32
	DiscLife             [4]float32
	IgnitionOn           int32
	StarterEngineOn      int32
	IsEngineRunning      int32
	KerbVibration        float32
	SlipVibration        float32
	GVibration           float32
	AbsVibration         float32
}

type pageGraphics struct {
	PacketID                 int32
	AccStatus                accStatusType
	AccSessionType           accSessionType
	CurrentTime              [15]uint16
	LastTime                 [15]uint16
	BestTime                 [15]uint16
	Split                    [15]uint16
	CompletedLaps            int32
	Position                 int32
	ICurrentTime             int32
	ILastTime                int32
	IBestTime                int32
	SessionTimeLeft          float32
	DistanceTravelled        float32
	IsInPit                  int32
	CurrentSectorIndex       int32
	LastSectorTime           int32
	NumberOfLaps             int32
	TyreCompound             [33]uint16
	_                        float32 // Unused replayTimeMultiplayer
	NormalizedCarPosition    float32
	ActiveCars               int32
	CarCordinates            [60][3]float32
	CarID                    [60]int32
	PlayerCarID              int32
	PenaltyTime              float32
	AccFlagType              accFlagType
	AccPenaltyType           accPenaltyType
	IdealLineOn              int32
	IsInPitLane              int32
	SurfaceGrip              float32
	MandatoryPitDone         int32
	WindSpeed                float32
	WindDirection            float32
	IsSetupMenuVisible       int32
	MainDisplayIndex         int32
	SecondaryDisplayIndex    int32
	TractionControl          int32
	TractionContolCut        int32
	EngineMap                int32
	Abs                      int32
	FuelXLap                 float32
	RainLights               int32
	FlashingLights           int32
	LightsStage              int32
	ExhaustTemperature       float32
	WiperLV                  int32
	DriverStintTotalTimeLeft int32
	DriverStintTimeLeft      int32
	RainTyres                int32
	SessionIndex             int32
	UsedFuel                 float32
	DeltaLapTime             [15]uint16
	IDeltaLapTime            int32
	EstimatedLapTime         [15]uint16
	IEstimatedLapTime        int32
	IsDeltaPositive          int32
	ISplit                   int32
	IsValidLap               int32
	FuelEstimatedLaps        float32
	TrackStatus              [33]uint16
	MissingMandatoryPits     int32
	Clock                    float32
	DirectionLightsLeft      int32
	DirectionLightsRight     int32
	GlobalYellow             int32
	GlobalYellow1            int32
	GlobalYellow2            int32
	GlobalYellow3            int32
	GlobalWhite              int32
	GlobalGreen              int32
	GlobalChequered          int32
	GlobalRed                int32
	MfdTyreSet               int32
	MfdFuelToAdd             float32
	MfdTyrePressureLF        float32
	MfdTyrePressureRF        float32
	MfdTyrePressureLR        float32
	MfdTyrePressureRR        float32
	AccTrackGripStatus       accTrackGripStatusType
	AccRainIntensity         accRainIntensityType
	AccRainIntensityIn10min  accRainIntensityType
	AccRainIntensityIn30min  accRainIntensityType
	CurrentTyreSet           int32
	StrategyTyreSet          int32
	GapAhead                 int32
	GapBehind                int32
}
