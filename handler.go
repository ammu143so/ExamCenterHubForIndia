package handler

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ExamCenterHandler handles all exam center assignment operations
type ExamCenterHandler struct {
	cities         map[string]City
	examCenters    map[string][]ExamCenter
	centerCapacity map[string]CenterCapacity
	registrations  []ExamRegistration
}

// StudentInfo holds user-provided student data for a run
type StudentInfo struct {
	Name       string
	ExamType   string
	RollNumber string
}

// CityDistance ties a city, distance and its centers
type CityDistance struct {
	City     City
	Distance float64
	Centers  []ExamCenter
}

// NewExamCenterHandler creates a new instance of ExamCenterHandler
func NewExamCenterHandler() *ExamCenterHandler {
	h := &ExamCenterHandler{
		cities:         make(map[string]City),
		examCenters:    make(map[string][]ExamCenter),
		centerCapacity: make(map[string]CenterCapacity),
		registrations:  make([]ExamRegistration, 0),
	}

	h.initializeCities()
	h.initializeExamCenters()
	h.initializeCenterCapacity()
	return h
}

// initializeCities populates the cities map with Indian cities and their coordinates
func (h *ExamCenterHandler) initializeCities() {
	h.cities = map[string]City{
		"Mumbai":     {Name: "Mumbai", Lat: 19.0760, Lng: 72.8777},
		"Delhi":      {Name: "Delhi", Lat: 28.7041, Lng: 77.1025},
		"Bangalore":  {Name: "Bangalore", Lat: 12.9716, Lng: 77.5946},
		"Hyderabad":  {Name: "Hyderabad", Lat: 17.3850, Lng: 78.4867},
		"Chennai":    {Name: "Chennai", Lat: 13.0827, Lng: 80.2707},
		"Kolkata":    {Name: "Kolkata", Lat: 22.5726, Lng: 88.3639},
		"Pune":       {Name: "Pune", Lat: 18.5204, Lng: 73.8567},
		"Ahmedabad":  {Name: "Ahmedabad", Lat: 23.0225, Lng: 72.5714},
		"Jaipur":     {Name: "Jaipur", Lat: 26.9124, Lng: 75.7873},
		"Lucknow":    {Name: "Lucknow", Lat: 26.8467, Lng: 80.9462},
		"Kanpur":     {Name: "Kanpur", Lat: 26.4499, Lng: 80.3319},
		"Nagpur":     {Name: "Nagpur", Lat: 21.1458, Lng: 79.0882},
		"Indore":     {Name: "Indore", Lat: 22.7196, Lng: 75.8577},
		"Bhopal":     {Name: "Bhopal", Lat: 23.2599, Lng: 77.4126},
		"Patna":      {Name: "Patna", Lat: 25.5941, Lng: 85.1376},
		"Vadodara":   {Name: "Vadodara", Lat: 22.3072, Lng: 73.1812},
		"Ghaziabad":  {Name: "Ghaziabad", Lat: 28.6692, Lng: 77.4538},
		"Agra":       {Name: "Agra", Lat: 27.1767, Lng: 78.0081},
		"Nashik":     {Name: "Nashik", Lat: 19.9975, Lng: 73.7898},
		"Faridabad":  {Name: "Faridabad", Lat: 28.4089, Lng: 77.3178},
		"Meerut":     {Name: "Meerut", Lat: 28.9845, Lng: 77.7064},
		"Rajkot":     {Name: "Rajkot", Lat: 22.3039, Lng: 70.8022},
		"Kalyan":     {Name: "Kalyan", Lat: 19.2403, Lng: 73.1305},
		"Vasai":      {Name: "Vasai", Lat: 19.4909, Lng: 72.8147},
		"Varanasi":   {Name: "Varanasi", Lat: 25.3176, Lng: 82.9739},
		"Srinagar":   {Name: "Srinagar", Lat: 34.0837, Lng: 74.7973},
		"Aurangabad": {Name: "Aurangabad", Lat: 19.8762, Lng: 75.3433},
		"Dhanbad":    {Name: "Dhanbad", Lat: 23.7957, Lng: 86.4304},
		"Amritsar":   {Name: "Amritsar", Lat: 31.6340, Lng: 74.8723},
		"Navi Mumbai": {Name: "Navi Mumbai", Lat: 19.0330, Lng: 73.0297},
		"Allahabad":  {Name: "Allahabad", Lat: 25.4358, Lng: 81.8463},
		"Ranchi":     {Name: "Ranchi", Lat: 23.3441, Lng: 85.3096},
		"Howrah":     {Name: "Howrah", Lat: 22.5958, Lng: 88.2636},
		"Coimbatore": {Name: "Coimbatore", Lat: 11.0168, Lng: 76.9558},
		"Jabalpur":   {Name: "Jabalpur", Lat: 23.1815, Lng: 79.9864},
		"Gwalior":    {Name: "Gwalior", Lat: 26.2183, Lng: 78.1828},
		"Vijayawada": {Name: "Vijayawada", Lat: 16.5062, Lng: 80.6480},
		"Jodhpur":    {Name: "Jodhpur", Lat: 26.2389, Lng: 73.0243},
		"Madurai":    {Name: "Madurai", Lat: 9.9252, Lng: 78.1198},
		"Raipur":     {Name: "Raipur", Lat: 21.2514, Lng: 81.6296},
		"Kota":       {Name: "Kota", Lat: 25.2138, Lng: 75.8648},
		"Chandigarh": {Name: "Chandigarh", Lat: 30.7333, Lng: 76.7794},
		"Guwahati":   {Name: "Guwahati", Lat: 26.1445, Lng: 91.7362},
		"Solapur":    {Name: "Solapur", Lat: 17.6599, Lng: 75.9064},
		"Hubli":      {Name: "Hubli", Lat: 15.3647, Lng: 75.1240},
		"Bareilly":   {Name: "Bareilly", Lat: 28.3670, Lng: 79.4304},
		"Moradabad":  {Name: "Moradabad", Lat: 28.8386, Lng: 78.7733},
		"Mysore":     {Name: "Mysore", Lat: 12.2958, Lng: 76.6394},
		"Gurgaon":    {Name: "Gurgaon", Lat: 28.4595, Lng: 77.0266},
		"Aligarh":    {Name: "Aligarh", Lat: 27.8974, Lng: 78.0880},
		"Jalandhar":  {Name: "Jalandhar", Lat: 31.3260, Lng: 75.5762},
	}
}

// initializeExamCenters populates the exam centers map
func (h *ExamCenterHandler) initializeExamCenters() {
	h.examCenters = map[string][]ExamCenter{
		"Mumbai": {
			{Name: "Mumbai Central Exam Center", City: "Mumbai"},
			{Name: "Bandra Kurla Complex Center", City: "Mumbai"},
			{Name: "Andheri Sports Complex", City: "Mumbai"},
		},
		"Delhi": {
			{Name: "Pragati Maidan Convention Center", City: "Delhi"},
			{Name: "Delhi University Exam Center", City: "Delhi"},
			{Name: "JNU Examination Hall", City: "Delhi"},
		},
		"Bangalore": {
			{Name: "Bangalore Palace Grounds", City: "Bangalore"},
			{Name: "IISC Exam Center", City: "Bangalore"},
			{Name: "Vidhana Soudha Center", City: "Bangalore"},
		},
		"Hyderabad": {
			{Name: "HITEC City Exam Center", City: "Hyderabad"},
			{Name: "University of Hyderabad Center", City: "Hyderabad"},
			{Name: "Gachibowli Stadium Center", City: "Hyderabad"},
		},
		"Chennai": {
			{Name: "Anna University Center", City: "Chennai"},
			{Name: "IIT Madras Exam Hall", City: "Chennai"},
			{Name: "Marina Beach Convention Center", City: "Chennai"},
		},
		"Kolkata": {
			{Name: "Salt Lake Stadium Center", City: "Kolkata"},
			{Name: "University of Calcutta Hall", City: "Kolkata"},
			{Name: "Science City Exam Center", City: "Kolkata"},
		},
		"Pune": {
			{Name: "Pune University Center", City: "Pune"},
			{Name: "Shivaji Nagar Exam Hall", City: "Pune"},
			{Name: "Kothrud Sports Complex", City: "Pune"},
		},
		"Ahmedabad": {
			{Name: "Gujarat University Center", City: "Ahmedabad"},
			{Name: "Sardar Patel Stadium Center", City: "Ahmedabad"},
			{Name: "IIM Ahmedabad Hall", City: "Ahmedabad"},
		},
		"Jaipur": {
			{Name: "Rajasthan University Center", City: "Jaipur"},
			{Name: "SMS Stadium Exam Hall", City: "Jaipur"},
			{Name: "Albert Hall Convention Center", City: "Jaipur"},
		},
		"Lucknow": {
			{Name: "Lucknow University Center", City: "Lucknow"},
			{Name: "Ekana Cricket Stadium Hall", City: "Lucknow"},
			{Name: "Gomti Riverfront Center", City: "Lucknow"},
		},
	}

	// Add default centers for remaining cities
	for cityName := range h.cities {
		if _, exists := h.examCenters[cityName]; !exists {
			h.examCenters[cityName] = []ExamCenter{
				{Name: fmt.Sprintf("%s Central Exam Center", cityName), City: cityName},
				{Name: fmt.Sprintf("%s University Center", cityName), City: cityName},
			}
		}
	}
}

// initializeCenterCapacity initializes capacity for all centers
func (h *ExamCenterHandler) initializeCenterCapacity() {
	for _, centers := range h.examCenters {
		for _, center := range centers {
			totalSeats := 200 + (len(center.Name) % 300)
			h.centerCapacity[center.Name] = CenterCapacity{
				TotalSeats:     totalSeats,
				AvailableSeats: totalSeats - (totalSeats / 10),
				BookedSeats:    totalSeats / 10,
			}
		}
	}
}

// GetAvailableCities returns a sorted list of all available cities
func (h *ExamCenterHandler) GetAvailableCities() []string {
	cities := make([]string, 0, len(h.cities))
	for name := range h.cities {
		cities = append(cities, name)
	}
	sort.Strings(cities)
	return cities
}

// ValidateCity checks if the given city exists in the system
func (h *ExamCenterHandler) ValidateCity(cityInput string) (string, error) {
	cities := h.GetAvailableCities()
	if cityInput == "" {
		return "", fmt.Errorf("city input cannot be empty")
	}
	if cityNumber, err := strconv.Atoi(cityInput); err == nil {
		if cityNumber < 1 || cityNumber > len(cities) {
			return "", fmt.Errorf("invalid city number")
		}
		return cities[cityNumber-1], nil
	}
	for _, cityName := range cities {
		if strings.EqualFold(cityName, cityInput) {
			return cityName, nil
		}
	}
	return "", fmt.Errorf("city '%s' not found in our database", cityInput)
}

// ValidateStudentInfo validates and returns student information
func (h *ExamCenterHandler) ValidateStudentInfo(name, examType, rollNumber string) (StudentInfo, error) {
	var student StudentInfo
	name = strings.TrimSpace(name)
	if name == "" {
		return student, fmt.Errorf("name cannot be empty")
	}
	examType = strings.TrimSpace(examType)
	if examType == "" {
		return student, fmt.Errorf("exam type cannot be empty")
	}
	rollNumber = strings.TrimSpace(rollNumber)
	if rollNumber == "" {
		return student, fmt.Errorf("roll number cannot be empty")
	}
	return StudentInfo{Name: name, ExamType: examType, RollNumber: rollNumber}, nil
}

// FindNearestCities finds nearest cities to the home city (excluding home city)
func (h *ExamCenterHandler) FindNearestCities(homeCity string, count int) ([]CityDistance, error) {
	homeCityData, exists := h.cities[homeCity]
	if !exists {
		return nil, fmt.Errorf("home city '%s' not found", homeCity)
	}
	var distances []CityDistance
	for cityName, cityData := range h.cities {
		if strings.EqualFold(cityName, homeCity) {
			continue
		}
		distance := h.calculateDistance(homeCityData, cityData)
		cityDistance := CityDistance{City: cityData, Distance: distance, Centers: h.examCenters[cityName]}
		distances = append(distances, cityDistance)
	}
	sort.Slice(distances, func(i, j int) bool { return distances[i].Distance < distances[j].Distance })
	if len(distances) > count {
		distances = distances[:count]
	}
	return distances, nil
}

// Advanced: find nearest applying preferences and capacity
func (h *ExamCenterHandler) FindNearestCitiesAdvanced(homeCity string, examType ExamType, preferences StudentPreference) ([]CityDistance, error) {
	homeCityData, exists := h.cities[homeCity]
	if !exists {
		return nil, fmt.Errorf("home city '%s' not found", homeCity)
	}
	var distances []CityDistance
	for cityName, cityData := range h.cities {
		if strings.EqualFold(cityName, homeCity) {
			continue
		}
		distance := h.calculateDistance(homeCityData, cityData)
		if preferences.MaxDistance > 0 && distance > preferences.MaxDistance {
			continue
		}
		available := h.getAvailableCenters(cityName)
		if len(available) == 0 {
			continue
		}
		distances = append(distances, CityDistance{City: cityData, Distance: distance, Centers: available})
	}
	sort.Slice(distances, func(i, j int) bool { return distances[i].Distance < distances[j].Distance })
	if max := examType.MaxCenters; max > 0 && len(distances) > max {
		distances = distances[:max]
	}
	return distances, nil
}

// getAvailableCenters returns centers with available seats
func (h *ExamCenterHandler) getAvailableCenters(cityName string) []ExamCenter {
	centers := h.examCenters[cityName]
	var available []ExamCenter
	for _, c := range centers {
		if capInfo, ok := h.centerCapacity[c.Name]; ok {
			if capInfo.AvailableSeats > 0 {
				available = append(available, c)
			}
		} else {
			available = append(available, c)
		}
	}
	return available
}

// calculateDistance calculates the distance between two cities using Haversine formula
func (h *ExamCenterHandler) calculateDistance(city1, city2 City) float64 {
	const earthRadius = 6371.0
	lat1 := h.toRadians(city1.Lat)
	lat2 := h.toRadians(city2.Lat)
	dLat := h.toRadians(city2.Lat - city1.Lat)
	dLon := h.toRadians(city2.Lng - city1.Lng)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadius * c
}

func (h *ExamCenterHandler) toRadians(deg float64) float64 { return deg * math.Pi / 180 }

// IO helpers
func (h *ExamCenterHandler) GetUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func (h *ExamCenterHandler) DisplayCityList() {
	fmt.Println("Available cities in India:")
	cities := h.GetAvailableCities()
	for i, city := range cities {
		fmt.Printf("%02d. %s\n", i+1, city)
	}
	fmt.Println()
}

// Basic flow
func (h *ExamCenterHandler) ProcessExamCenterAssignment() error {
	h.DisplayCityList()
	cityInput, err := h.GetUserInput("Enter your home city (name or number): ")
	if err != nil { return fmt.Errorf("error reading city input: %v", err) }
	homeCity, err := h.ValidateCity(cityInput)
	if err != nil { return err }
	name, err := h.GetUserInput("Enter your name: ")
	if err != nil { return fmt.Errorf("error reading name: %v", err) }
	examType, err := h.GetUserInput("Enter exam type (e.g., JEE, NEET, UPSC, etc.): ")
	if err != nil { return fmt.Errorf("error reading exam type: %v", err) }
	roll, err := h.GetUserInput("Enter your roll number/application number: ")
	if err != nil { return fmt.Errorf("error reading roll number: %v", err) }
	student, err := h.ValidateStudentInfo(name, examType, roll)
	if err != nil { return err }
	nearest, err := h.FindNearestCities(homeCity, 3)
	if err != nil { return err }
	h.DisplayResults(student, homeCity, nearest)
	return nil
}

func (h *ExamCenterHandler) DisplayResults(student StudentInfo, homeCity string, nearest []CityDistance) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("EXAMINATION CENTER ASSIGNMENT RESULT")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Student Name: %s\n", student.Name)
	fmt.Printf("Roll Number: %s\n", student.RollNumber)
	fmt.Printf("Exam Type: %s\n", student.ExamType)
	fmt.Printf("Home City: %s\n", homeCity)
	fmt.Printf("Country: India\n")
	fmt.Println("\nASSIGNED EXAMINATION CENTERS (Nearest cities excluding home city):")
	fmt.Println(strings.Repeat("-", 60))
	for i, cd := range nearest {
		fmt.Printf("\n%d. %s\n", i+1, strings.ToUpper(cd.City.Name))
		fmt.Printf("   Distance from %s: %.1f km\n", homeCity, cd.Distance)
		fmt.Printf("   Available Centers:\n")
		for _, center := range cd.Centers {
			fmt.Printf("   • %s\n", center.Name)
		}
	}
	if len(nearest) > 0 && len(nearest[0].Centers) > 0 {
		fmt.Printf("\nRECOMMENDED CENTER: %s\n", nearest[0].Centers[0].Name)
		fmt.Printf("Location: %s (%.1f km from your home city)\n", nearest[0].City.Name, nearest[0].Distance)
	}
}

// Advanced flow
func (h *ExamCenterHandler) DisplayExamTypes() {
	fmt.Println("Available Exam Types:")
	fmt.Println("=====================")
	for code, exam := range PredefinedExamTypes {
		fmt.Printf("%s - %s\n", code, exam.Name)
		fmt.Printf("   Duration: %s | Max Centers: %d\n", exam.Duration.String(), exam.MaxCenters)
		fmt.Printf("   %s\n\n", exam.Description)
	}
}

func (h *ExamCenterHandler) GetExamTypeDetails(examCode string) (ExamType, error) {
	ex, ok := PredefinedExamTypes[strings.ToUpper(examCode)]
	if !ok { return ExamType{}, fmt.Errorf("exam type '%s' not found", examCode) }
	return ex, nil
}

func (h *ExamCenterHandler) ProcessAdvancedExamAssignment() error {
	fmt.Println("=== Advanced Exam Center Assignment ===")
	h.DisplayExamTypes()
	examInput, err := h.GetUserInput("Enter exam type (e.g., JEE, NEET, UPSC): ")
	if err != nil { return fmt.Errorf("error reading exam type: %v", err) }
	exType, err := h.GetExamTypeDetails(examInput)
	if err != nil { return err }
	fmt.Printf("\nSelected: %s - %s\n", exType.Code, exType.Name)
	fmt.Printf("Duration: %s\n\n", exType.Duration.String())
	h.DisplayCityList()
	cityInput, err := h.GetUserInput("Enter your home city (name or number): ")
	if err != nil { return fmt.Errorf("error reading city input: %v", err) }
	homeCity, err := h.ValidateCity(cityInput)
	if err != nil { return err }
	name, err := h.GetUserInput("Enter your name: ")
	if err != nil { return fmt.Errorf("error reading name: %v", err) }
	roll, err := h.GetUserInput("Enter your roll number/application number: ")
	if err != nil { return fmt.Errorf("error reading roll number: %v", err) }
	student, err := h.ValidateStudentInfo(name, exType.Code, roll)
	if err != nil { return err }
	prefs, err := h.GetStudentPreferences()
	if err != nil { return err }
	nearest, err := h.FindNearestCitiesAdvanced(homeCity, exType, prefs)
	if err != nil { return err }
	if len(nearest) == 0 { return fmt.Errorf("no suitable exam centers found within your preferences") }
	reg := h.CreateRegistration(student, exType, nearest[0], homeCity)
	h.DisplayAdvancedResults(reg, nearest, prefs)
	return nil
}

func (h *ExamCenterHandler) GetStudentPreferences() (StudentPreference, error) {
	var p StudentPreference
	maxDist, err := h.GetUserInput("Maximum acceptable distance (km) [default: 1000]: ")
	if err != nil { return p, err }
	if maxDist == "" {
		p.MaxDistance = 1000.0
	} else if v, err := strconv.ParseFloat(maxDist, 64); err == nil {
		p.MaxDistance = v
	} else {
		p.MaxDistance = 1000.0
	}
	transport, err := h.GetUserInput("Preferred transport mode (train/flight/bus) [default: any]: ")
	if err != nil { return p, err }
	p.PreferredTransport = strings.TrimSpace(transport)
	acc, err := h.GetUserInput("Need accommodation? (y/n) [default: n]: ")
	if err != nil { return p, err }
	p.AccommodationNeeded = strings.ToLower(acc) == "y" || strings.ToLower(acc) == "yes"
	return p, nil
}

// Registration helpers
func (h *ExamCenterHandler) CreateRegistration(student StudentInfo, examType ExamType, assigned CityDistance, homeCity string) ExamRegistration {
	reg := ExamRegistration{
		ID:               h.generateRegistrationID(examType.Code, student.RollNumber),
		StudentName:      student.Name,
		StudentCity:      homeCity,
		ExamType:         examType,
		AssignedCenter:   assigned.Centers[0].Name,
		AssignedCity:     assigned.City.Name,
		Distance:         assigned.Distance,
		RegistrationTime: time.Now(),
		Preferences:      StudentPreference{MaxDistance: 1000, PreferredTransport: "any", AccommodationNeeded: false},
	}
	h.registrations = append(h.registrations, reg)
	if capInfo, ok := h.centerCapacity[reg.AssignedCenter]; ok {
		capInfo.AvailableSeats--
		capInfo.BookedSeats++
		h.centerCapacity[reg.AssignedCenter] = capInfo
	}
	return reg
}

func (h *ExamCenterHandler) generateRegistrationID(examCode, roll string) string {
	return fmt.Sprintf("%s-%s-%s", examCode, roll, time.Now().Format("20060102150405"))
}

func (h *ExamCenterHandler) DisplayAdvancedResults(reg ExamRegistration, nearest []CityDistance, prefs StudentPreference) {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("ADVANCED EXAMINATION CENTER ASSIGNMENT RESULT")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("Registration ID: %s\n", reg.ID)
	fmt.Printf("Student Name: %s\n", reg.StudentName)
	fmt.Printf("Exam: %s - %s\n", reg.ExamType.Code, reg.ExamType.Name)
	fmt.Printf("Duration: %s\n", reg.ExamType.Duration.String())
	fmt.Printf("Registered: %s\n", reg.RegistrationTime.Format("2006-01-02 15:04:05"))
	fmt.Println("\n" + strings.Repeat("-", 70))
	fmt.Println("ASSIGNED CENTER:")
	fmt.Printf("🏢 Center: %s\n", reg.AssignedCenter)
	fmt.Printf("🏙️  City: %s\n", reg.AssignedCity)
	fmt.Printf("📏 Distance: %.1f km from your home city\n", reg.Distance)
	if capInfo, ok := h.centerCapacity[reg.AssignedCenter]; ok {
		fmt.Printf("💺 Capacity: %d total, %d available, %d booked\n", capInfo.TotalSeats, capInfo.AvailableSeats, capInfo.BookedSeats)
	}
	fmt.Println("\n" + strings.Repeat("-", 70))
	fmt.Println("ALTERNATIVE OPTIONS:")
	for i, cd := range nearest {
		if i == 0 { continue }
		fmt.Printf("\n%d. %s (%.1f km)\n", i+1, strings.ToUpper(cd.City.Name), cd.Distance)
		for _, c := range cd.Centers {
			if capInfo, ok := h.centerCapacity[c.Name]; ok {
				fmt.Printf("   • %s [%d seats available]\n", c.Name, capInfo.AvailableSeats)
			} else {
				fmt.Printf("   • %s\n", c.Name)
			}
		}
	}
	fmt.Println("\n" + strings.Repeat("-", 70))
	fmt.Println("STUDENT PREFERENCES APPLIED:")
	fmt.Printf("• Max Distance: %.0f km\n", prefs.MaxDistance)
	if prefs.PreferredTransport != "" {
		fmt.Printf("• Transport Mode: %s\n", prefs.PreferredTransport)
	}
	if prefs.AccommodationNeeded {
		fmt.Println("• Accommodation: Required")
	}
	fmt.Println("\n📋 IMPORTANT INSTRUCTIONS:")
	fmt.Println("• Save your Registration ID for future reference")
	fmt.Println("• Carry this assignment along with your admit card")
	fmt.Println("• Reach the center at least 30 minutes before exam time")
	fmt.Printf("• Exam duration: %s\n", reg.ExamType.Duration.String())
}

func (h *ExamCenterHandler) ShowRegistrationSummary() {
	if len(h.registrations) == 0 {
		fmt.Println("No registrations found.")
		return
	}
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("REGISTRATION SUMMARY")
	fmt.Println(strings.Repeat("=", 60))
	for i, reg := range h.registrations {
		fmt.Printf("\n%d. %s (%s)\n", i+1, reg.StudentName, reg.ExamType.Code)
		fmt.Printf("   ID: %s\n", reg.ID)
		fmt.Printf("   Center: %s, %s\n", reg.AssignedCenter, reg.AssignedCity)
		fmt.Printf("   Distance: %.1f km\n", reg.Distance)
	}
} 