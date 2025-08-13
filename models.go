package handler

import "time"

// City represents a city with its coordinates
type City struct {
	Name string
	Lat  float64
	Lng  float64
}

// ExamCenter represents an examination center
type ExamCenter struct {
	Name string
	City string
}

// ExamType represents different types of examinations
type ExamType struct {
	Code        string
	Name        string
	Description string
	Duration    time.Duration
	Schedule    ExamSchedule
	MaxCenters  int // Max number of nearby cities to suggest
}

// ExamSchedule represents the schedule information for an exam
type ExamSchedule struct {
	StartDate            string
	EndDate              string
	TimeSlots            []string
	RegistrationDeadline string
}

// CenterCapacity represents the capacity information for each center
type CenterCapacity struct {
	TotalSeats     int
	AvailableSeats int
	BookedSeats    int
}

// StudentPreference represents student's preferences for exam assignment
type StudentPreference struct {
	MaxDistance        float64 // km
	PreferredTransport string  // "train" | "bus" | "flight" | "any"
	AccommodationNeeded bool
}

// ExamRegistration represents a completed exam registration
type ExamRegistration struct {
	ID               string
	StudentName      string
	StudentCity      string
	ExamType         ExamType
	AssignedCity     string
	AssignedCenter   string
	Distance         float64
	RegistrationTime time.Time
	Preferences      StudentPreference
}

// PredefinedExamTypes contains commonly available exam types in India
var PredefinedExamTypes = map[string]ExamType{
	"JEE": {
		Code:        "JEE",
		Name:        "Joint Entrance Examination",
		Description: "Engineering entrance exam for IITs, NITs, and other technical institutes",
		Duration:    3 * time.Hour,
		Schedule: ExamSchedule{
			StartDate:            "2024-04-01",
			EndDate:              "2024-04-30",
			TimeSlots:            []string{"09:00-12:00", "15:00-18:00"},
			RegistrationDeadline: "2024-03-15",
		},
		MaxCenters: 3,
	},
	"NEET": {
		Code:        "NEET",
		Name:        "National Eligibility cum Entrance Test",
		Description: "Medical entrance exam for MBBS, BDS, and other medical courses",
		Duration:    3*time.Hour + 20*time.Minute,
		Schedule: ExamSchedule{
			StartDate:            "2024-05-05",
			EndDate:              "2024-05-05",
			TimeSlots:            []string{"14:00-17:20"},
			RegistrationDeadline: "2024-04-15",
		},
		MaxCenters: 2,
	},
	"UPSC": {
		Code:        "UPSC",
		Name:        "Union Public Service Commission",
		Description: "Civil services examination for IAS, IPS, IFS and other central services",
		Duration:    6 * time.Hour,
		Schedule: ExamSchedule{
			StartDate:            "2024-06-02",
			EndDate:              "2024-06-04",
			TimeSlots:            []string{"09:30-12:30", "14:30-17:30"},
			RegistrationDeadline: "2024-05-01",
		},
		MaxCenters: 2,
	},
	"CAT": {
		Code:        "CAT",
		Name:        "Common Admission Test",
		Description: "Management entrance exam for IIMs and other business schools",
		Duration:    2*time.Hour + 40*time.Minute,
		Schedule: ExamSchedule{
			StartDate:            "2024-11-26",
			EndDate:              "2024-11-26",
			TimeSlots:            []string{"08:30-11:10", "14:30-17:10", "18:30-21:10"},
			RegistrationDeadline: "2024-09-20",
		},
		MaxCenters: 4,
	},
	"GATE": {
		Code:        "GATE",
		Name:        "Graduate Aptitude Test in Engineering",
		Description: "Entrance exam for M.Tech, PhD and PSU recruitments",
		Duration:    3 * time.Hour,
		Schedule: ExamSchedule{
			StartDate:            "2024-02-03",
			EndDate:              "2024-02-11",
			TimeSlots:            []string{"09:30-12:30", "14:30-17:30"},
			RegistrationDeadline: "2024-01-03",
		},
		MaxCenters: 3,
	},
	"SSC": {
		Code:        "SSC",
		Name:        "Staff Selection Commission",
		Description: "Recruitment exam for various government departments",
		Duration:    2 * time.Hour,
		Schedule: ExamSchedule{
			StartDate:            "2024-07-01",
			EndDate:              "2024-07-25",
			TimeSlots:            []string{"10:00-12:00", "14:30-16:30"},
			RegistrationDeadline: "2024-06-01",
		},
		MaxCenters: 5,
	},
	"IBPS": {
		Code:        "IBPS",
		Name:        "Institute of Banking Personnel Selection",
		Description: "Banking sector recruitment examination",
		Duration:    2*time.Hour + 45*time.Minute,
		Schedule: ExamSchedule{
			StartDate:            "2024-08-17",
			EndDate:              "2024-08-25",
			TimeSlots:            []string{"09:00-11:45", "13:30-16:15"},
			RegistrationDeadline: "2024-07-15",
		},
		MaxCenters: 4,
	},
	"IELTS": {
		Code:        "IELTS",
		Name:        "International English Language Testing System",
		Description: "English proficiency test for international education and migration",
		Duration:    2*time.Hour + 45*time.Minute,
		Schedule: ExamSchedule{
			StartDate:            "2024-01-01",
			EndDate:              "2024-12-31",
			TimeSlots:            []string{"09:00-12:00", "13:00-16:00"},
			RegistrationDeadline: "Rolling basis",
		},
		MaxCenters: 3,
	},
} 