package phishing

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// PhishingCampaign represents a phishing campaign
type PhishingCampaign struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Targets     []string  `json:"targets"`
	Template    string    `json:"template"`
	Payload     string    `json:"payload"`
	Status      string    `json:"status"`
	Created     time.Time `json:"created"`
	Completed   time.Time `json:"completed,omitempty"`
	ClickRate   float64   `json:"click_rate"`
	SuccessRate float64   `json:"success_rate"`
}

// PhishingGenerator generates phishing emails
type PhishingGenerator struct {
	OutputDir string
}

// NewPhishingGenerator creates a new generator
func NewPhishingGenerator(outputDir string) *PhishingGenerator {
	return &PhishingGenerator{
		OutputDir: outputDir,
	}
}

// GenerateCampaign generates a phishing campaign
func (g *PhishingGenerator) GenerateCampaign(name string, targets []string, template string) (*PhishingCampaign, error) {
	fmt.Printf("Generating phishing campaign: %s\n", name)

	campaign := &PhishingCampaign{
		ID:       fmt.Sprintf("camp-%d", time.Now().Unix()),
		Name:     name,
		Targets:  targets,
		Template: template,
		Status:   "generated",
		Created:  time.Now(),
	}

	// Generate individual emails
	campaignDir := filepath.Join(g.OutputDir, "campaigns", campaign.ID)
	if err := os.MkdirAll(campaignDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create campaign directory: %w", err)
	}

	for _, target := range targets {
		email := g.generateEmail(target, template)
		emailPath := filepath.Join(campaignDir, fmt.Sprintf("%s.html", target))
		if err := os.WriteFile(emailPath, []byte(email), 0644); err != nil {
			return nil, fmt.Errorf("failed to write email: %w", err)
		}
	}

	// Save campaign metadata
	if err := g.saveCampaign(campaign); err != nil {
		return nil, fmt.Errorf("failed to save campaign: %w", err)
	}

	fmt.Printf("Campaign generated successfully: %s\n", campaign.ID)
	return campaign, nil
}

// generateEmail generates a single phishing email
func (g *PhishingGenerator) generateEmail(target string, template string) string {
	// In a real implementation, this would:
	// - Parse the template
	// - Insert personalized data (name, company, etc.)
	// - Generate tracking pixels and links
	// - Add HTML formatting and branding

	return fmt.Sprintf(`
<html>
<head><title>Important Notice</title></head>
<body>
<h2>Dear %s,</h2>
<p>We have detected unusual activity on your account. Please click the link below to verify:</p>
<a href="https://tracking.nightfall.local/verify?id=%s">Verify Now</a>
<p>Best regards,<br>Account Security Team</p>
</body>
</html>
`, target, target)
}

// saveCampaign saves campaign metadata
func (g *PhishingGenerator) saveCampaign(campaign *PhishingCampaign) error {
	data, err := json.MarshalIndent(campaign, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal campaign: %w", err)
	}

	campaignsDir := filepath.Join(g.OutputDir, "campaigns")
	if err := os.MkdirAll(campaignsDir, 0755); err != nil {
		return fmt.Errorf("failed to create campaigns directory: %w", err)
	}

	path := filepath.Join(campaignsDir, fmt.Sprintf("%s.json", campaign.ID))
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write campaign file: %w", err)
	}

	return nil
}
