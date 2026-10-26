package telemetry

type accStatusType int32

const (
	statusOff    accStatusType = 0
	statusReplay accStatusType = 1
	statusLive   accStatusType = 2
	statusPause  accStatusType = 3
)

type accFlagType int32

const (
	flagNone      accFlagType = 0
	flagBlue      accFlagType = 1
	flagYellow    accFlagType = 2
	flagBlack     accFlagType = 3
	flagWhite     accFlagType = 4
	flagCheckered accFlagType = 5
	flagPenalty   accFlagType = 6
	flagGreen     accFlagType = 7
	flagOrange    accFlagType = 8
)

type accPenaltyType int32

const (
	penaltyNone                             accPenaltyType = 0
	penaltyCuttingDrivethrough              accPenaltyType = 1
	penaltyCuttingStopnGo10                 accPenaltyType = 2
	penaltyCuttingStopnGo20                 accPenaltyType = 3
	penaltyCuttingStopnGo30                 accPenaltyType = 4
	penaltyCuttingDisqualified              accPenaltyType = 5
	penaltyCuttingRemoveBestTime            accPenaltyType = 6
	penaltyPitLimitDrivethrough             accPenaltyType = 7
	penaltyPitLimitStopnGo10                accPenaltyType = 8
	penaltyPitLimitStopnGo20                accPenaltyType = 9
	penaltyPitLimitStopnGo30                accPenaltyType = 10
	penaltyPitLimitDisqualified             accPenaltyType = 11
	penaltyPitLimitRemoveBestTime           accPenaltyType = 12
	penaltyDisqualifiedMandatoryPitstop     accPenaltyType = 13
	penaltyPostRaceTime                     accPenaltyType = 14
	penaltyDisqualifiedTrolling             accPenaltyType = 15
	penaltyDisqualifiedPitEntry             accPenaltyType = 16
	penaltyDisqualifiedPitExit              accPenaltyType = 17
	penaltyDisqualifiedWrongWay             accPenaltyType = 18
	penaltyDrivethroughIgnoredDriverStint   accPenaltyType = 19
	penaltyDisqualifiedIgnoredDriverStint   accPenaltyType = 20
	penaltyDisqualifiedExceedingDriverStint accPenaltyType = 21
)

type accSessionType int32

const (
	sessionUnknow            accSessionType = -1
	sessionPractice          accSessionType = 0
	sessionQualy             accSessionType = 1
	sessionRace              accSessionType = 2
	sessionHotlap            accSessionType = 3
	sessionTimeAttack        accSessionType = 4
	sessionDrift             accSessionType = 5
	sessionDrag              accSessionType = 6
	sessionHotstint          accSessionType = 7
	sessionHotstintSuperPole accSessionType = 8
)

type accRainIntensityType int32

const (
	rainNone         accRainIntensityType = 0
	rainDrizzle      accRainIntensityType = 1
	rainLight        accRainIntensityType = 2
	rainMedium       accRainIntensityType = 3
	rainHeavy        accRainIntensityType = 4
	rainThunderstorm accRainIntensityType = 5
)

type accTrackGripStatusType int32

const (
	gripGreen   accTrackGripStatusType = 0
	gripFast    accTrackGripStatusType = 1
	gripOptimal accTrackGripStatusType = 2
	gripGreasy  accTrackGripStatusType = 3
	gripDamp    accTrackGripStatusType = 4
	gripWet     accTrackGripStatusType = 5
	gripFlooded accTrackGripStatusType = 6
)

type accWheelsType int32

const (
	wheelFrontLeft  accWheelsType = 0
	wheelFrontRight accWheelsType = 1
	wheelRearLeft   accWheelsType = 2
	wheelRearRight  accWheelsType = 3
)
