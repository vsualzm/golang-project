package database

// func LoadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file: %v", err)
// 	}
// }

// func ConnectDB() (*sql.DB, error) {
// 	// Load environment variables
// 	LoadEnv()

// 	fmt.Println("Cheking ENV variables")
// 	fmt.Println(os.Getenv("DB_HOST"))
// 	fmt.Println(os.Getenv("DB_PORT"))
// 	fmt.Println(os.Getenv("DB_USER"))
// 	fmt.Println(os.Getenv("DB_PASSWORD"))

// 	// Build connection string from .env variables
// 	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		os.Getenv("DB_HOST"),
// 		os.Getenv("DB_PORT"),
// 		os.Getenv("DB_USER"),
// 		os.Getenv("DB_PASSWORD"),
// 		os.Getenv("DB_NAME"))

// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open database: %w", err)
// 	}

// 	// Test the connection
// 	if err := db.Ping(); err != nil {
// 		return nil, fmt.Errorf("failed to connect to database: %w", err)
// 	}

// 	return db, nil
// }
