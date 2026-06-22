package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"http://localhost:5174",
	"https://belatihan-production.up.railway.app",
	"https://my-fe-rho.vercel.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}