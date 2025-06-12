package rules

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"iiot-backend/models"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

// Rule methods
func (s *Service) GetRules(enabledFilter *bool, limit, offset int) ([]models.Rule, error) {
	query := `
		SELECT id, name, description, enabled, priority, conditions, actions, tags, created, modified
		FROM rules
		WHERE ($1::boolean IS NULL OR enabled = $1)
		ORDER BY priority DESC, created DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := s.db.Query(query, enabledFilter, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query rules: %w", err)
	}
	defer rows.Close()

	var rules []models.Rule
	for rows.Next() {
		var rule models.Rule
		var conditionsJSON, actionsJSON, tagsJSON []byte

		err := rows.Scan(
			&rule.ID, &rule.Name, &rule.Description, &rule.Enabled, &rule.Priority,
			&conditionsJSON, &actionsJSON, &tagsJSON, &rule.Created, &rule.Modified,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rule: %w", err)
		}

		// Unmarshal JSON fields
		if len(conditionsJSON) > 0 {
			json.Unmarshal(conditionsJSON, &rule.Conditions)
		}
		if len(actionsJSON) > 0 {
			json.Unmarshal(actionsJSON, &rule.Actions)
		}
		if len(tagsJSON) > 0 {
			json.Unmarshal(tagsJSON, &rule.Tags)
		}

		rules = append(rules, rule)
	}

	return rules, nil
}

func (s *Service) GetRuleByID(id string) (*models.Rule, error) {
	query := `
		SELECT id, name, description, enabled, priority, conditions, actions, tags, created, modified
		FROM rules
		WHERE id = $1
	`

	var rule models.Rule
	var conditionsJSON, actionsJSON, tagsJSON []byte

	err := s.db.QueryRow(query, id).Scan(
		&rule.ID, &rule.Name, &rule.Description, &rule.Enabled, &rule.Priority,
		&conditionsJSON, &actionsJSON, &tagsJSON, &rule.Created, &rule.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get rule: %w", err)
	}

	// Unmarshal JSON fields
	if len(conditionsJSON) > 0 {
		json.Unmarshal(conditionsJSON, &rule.Conditions)
	}
	if len(actionsJSON) > 0 {
		json.Unmarshal(actionsJSON, &rule.Actions)
	}
	if len(tagsJSON) > 0 {
		json.Unmarshal(tagsJSON, &rule.Tags)
	}

	return &rule, nil
}

func (s *Service) GetRuleByName(name string) (*models.Rule, error) {
	query := `
		SELECT id, name, description, enabled, priority, conditions, actions, tags, created, modified
		FROM rules
		WHERE name = $1
	`

	var rule models.Rule
	var conditionsJSON, actionsJSON, tagsJSON []byte

	err := s.db.QueryRow(query, name).Scan(
		&rule.ID, &rule.Name, &rule.Description, &rule.Enabled, &rule.Priority,
		&conditionsJSON, &actionsJSON, &tagsJSON, &rule.Created, &rule.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get rule: %w", err)
	}

	// Unmarshal JSON fields
	if len(conditionsJSON) > 0 {
		json.Unmarshal(conditionsJSON, &rule.Conditions)
	}
	if len(actionsJSON) > 0 {
		json.Unmarshal(actionsJSON, &rule.Actions)
	}
	if len(tagsJSON) > 0 {
		json.Unmarshal(tagsJSON, &rule.Tags)
	}

	return &rule, nil
}

func (s *Service) CreateRule(req *models.RuleRequest) (string, error) {
	rule := &models.Rule{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		Enabled:     req.Enabled,
		Priority:    req.Priority,
		Conditions:  req.Conditions,
		Actions:     req.Actions,
		Tags:        req.Tags,
		Created:     time.Now(),
		Modified:    time.Now(),
	}

	// Marshal JSON fields
	conditionsJSON, _ := json.Marshal(rule.Conditions)
	actionsJSON, _ := json.Marshal(rule.Actions)
	tagsJSON, _ := json.Marshal(rule.Tags)

	query := `
		INSERT INTO rules (id, name, description, enabled, priority, conditions, actions, tags, created, modified)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := s.db.Exec(query, rule.ID, rule.Name, rule.Description, rule.Enabled,
		rule.Priority, conditionsJSON, actionsJSON, tagsJSON, rule.Created, rule.Modified)
	if err != nil {
		return "", fmt.Errorf("failed to create rule: %w", err)
	}

	return rule.ID, nil
}

func (s *Service) UpdateRule(id string, req *models.RuleRequest) error {
	// Marshal JSON fields
	conditionsJSON, _ := json.Marshal(req.Conditions)
	actionsJSON, _ := json.Marshal(req.Actions)
	tagsJSON, _ := json.Marshal(req.Tags)

	query := `
		UPDATE rules 
		SET name = $2, description = $3, enabled = $4, priority = $5, 
		    conditions = $6, actions = $7, tags = $8, modified = $9
		WHERE id = $1
	`

	_, err := s.db.Exec(query, id, req.Name, req.Description, req.Enabled, req.Priority,
		conditionsJSON, actionsJSON, tagsJSON, time.Now())
	if err != nil {
		return fmt.Errorf("failed to update rule: %w", err)
	}

	return nil
}

func (s *Service) DeleteRule(id string) error {
	query := `DELETE FROM rules WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete rule: %w", err)
	}
	return nil
}

func (s *Service) SetRuleEnabled(id string, enabled bool) error {
	query := `UPDATE rules SET enabled = $2, modified = $3 WHERE id = $1`
	_, err := s.db.Exec(query, id, enabled, time.Now())
	if err != nil {
		return fmt.Errorf("failed to update rule enabled status: %w", err)
	}
	return nil
}

// Pipeline methods
func (s *Service) GetPipelines(enabledFilter *bool, limit, offset int) ([]models.Pipeline, error) {
	query := `
		SELECT id, name, description, enabled, functions, triggers, targets, tags, created, modified
		FROM pipelines
		WHERE ($1::boolean IS NULL OR enabled = $1)
		ORDER BY created DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := s.db.Query(query, enabledFilter, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query pipelines: %w", err)
	}
	defer rows.Close()

	var pipelines []models.Pipeline
	for rows.Next() {
		var pipeline models.Pipeline
		var functionsJSON, triggersJSON, targetsJSON, tagsJSON []byte

		err := rows.Scan(
			&pipeline.ID, &pipeline.Name, &pipeline.Description, &pipeline.Enabled,
			&functionsJSON, &triggersJSON, &targetsJSON, &tagsJSON,
			&pipeline.Created, &pipeline.Modified,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan pipeline: %w", err)
		}

		// Unmarshal JSON fields
		if len(functionsJSON) > 0 {
			json.Unmarshal(functionsJSON, &pipeline.Functions)
		}
		if len(triggersJSON) > 0 {
			json.Unmarshal(triggersJSON, &pipeline.Triggers)
		}
		if len(targetsJSON) > 0 {
			json.Unmarshal(targetsJSON, &pipeline.Targets)
		}
		if len(tagsJSON) > 0 {
			json.Unmarshal(tagsJSON, &pipeline.Tags)
		}

		pipelines = append(pipelines, pipeline)
	}

	return pipelines, nil
}

func (s *Service) GetPipelineByID(id string) (*models.Pipeline, error) {
	query := `
		SELECT id, name, description, enabled, functions, triggers, targets, tags, created, modified
		FROM pipelines
		WHERE id = $1
	`

	var pipeline models.Pipeline
	var functionsJSON, triggersJSON, targetsJSON, tagsJSON []byte

	err := s.db.QueryRow(query, id).Scan(
		&pipeline.ID, &pipeline.Name, &pipeline.Description, &pipeline.Enabled,
		&functionsJSON, &triggersJSON, &targetsJSON, &tagsJSON,
		&pipeline.Created, &pipeline.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get pipeline: %w", err)
	}

	// Unmarshal JSON fields
	if len(functionsJSON) > 0 {
		json.Unmarshal(functionsJSON, &pipeline.Functions)
	}
	if len(triggersJSON) > 0 {
		json.Unmarshal(triggersJSON, &pipeline.Triggers)
	}
	if len(targetsJSON) > 0 {
		json.Unmarshal(targetsJSON, &pipeline.Targets)
	}
	if len(tagsJSON) > 0 {
		json.Unmarshal(tagsJSON, &pipeline.Tags)
	}

	return &pipeline, nil
}

func (s *Service) GetPipelineByName(name string) (*models.Pipeline, error) {
	query := `
		SELECT id, name, description, enabled, functions, triggers, targets, tags, created, modified
		FROM pipelines
		WHERE name = $1
	`

	var pipeline models.Pipeline
	var functionsJSON, triggersJSON, targetsJSON, tagsJSON []byte

	err := s.db.QueryRow(query, name).Scan(
		&pipeline.ID, &pipeline.Name, &pipeline.Description, &pipeline.Enabled,
		&functionsJSON, &triggersJSON, &targetsJSON, &tagsJSON,
		&pipeline.Created, &pipeline.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get pipeline: %w", err)
	}

	// Unmarshal JSON fields
	if len(functionsJSON) > 0 {
		json.Unmarshal(functionsJSON, &pipeline.Functions)
	}
	if len(triggersJSON) > 0 {
		json.Unmarshal(triggersJSON, &pipeline.Triggers)
	}
	if len(targetsJSON) > 0 {
		json.Unmarshal(targetsJSON, &pipeline.Targets)
	}
	if len(tagsJSON) > 0 {
		json.Unmarshal(tagsJSON, &pipeline.Tags)
	}

	return &pipeline, nil
}

func (s *Service) CreatePipeline(req *models.PipelineRequest) (string, error) {
	pipeline := &models.Pipeline{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		Enabled:     req.Enabled,
		Functions:   req.Functions,
		Triggers:    req.Triggers,
		Targets:     req.Targets,
		Tags:        req.Tags,
		Created:     time.Now(),
		Modified:    time.Now(),
	}

	// Marshal JSON fields
	functionsJSON, _ := json.Marshal(pipeline.Functions)
	triggersJSON, _ := json.Marshal(pipeline.Triggers)
	targetsJSON, _ := json.Marshal(pipeline.Targets)
	tagsJSON, _ := json.Marshal(pipeline.Tags)

	query := `
		INSERT INTO pipelines (id, name, description, enabled, functions, triggers, targets, tags, created, modified)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := s.db.Exec(query, pipeline.ID, pipeline.Name, pipeline.Description, pipeline.Enabled,
		functionsJSON, triggersJSON, targetsJSON, tagsJSON, pipeline.Created, pipeline.Modified)
	if err != nil {
		return "", fmt.Errorf("failed to create pipeline: %w", err)
	}

	return pipeline.ID, nil
}

func (s *Service) UpdatePipeline(id string, req *models.PipelineRequest) error {
	// Marshal JSON fields
	functionsJSON, _ := json.Marshal(req.Functions)
	triggersJSON, _ := json.Marshal(req.Triggers)
	targetsJSON, _ := json.Marshal(req.Targets)
	tagsJSON, _ := json.Marshal(req.Tags)

	query := `
		UPDATE pipelines 
		SET name = $2, description = $3, enabled = $4, functions = $5, 
		    triggers = $6, targets = $7, tags = $8, modified = $9
		WHERE id = $1
	`

	_, err := s.db.Exec(query, id, req.Name, req.Description, req.Enabled,
		functionsJSON, triggersJSON, targetsJSON, tagsJSON, time.Now())
	if err != nil {
		return fmt.Errorf("failed to update pipeline: %w", err)
	}

	return nil
}

func (s *Service) DeletePipeline(id string) error {
	query := `DELETE FROM pipelines WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete pipeline: %w", err)
	}
	return nil
}

func (s *Service) SetPipelineEnabled(id string, enabled bool) error {
	query := `UPDATE pipelines SET enabled = $2, modified = $3 WHERE id = $1`
	_, err := s.db.Exec(query, id, enabled, time.Now())
	if err != nil {
		return fmt.Errorf("failed to update pipeline enabled status: %w", err)
	}
	return nil
}

// Rule execution methods
func (s *Service) ExecuteRule(id string) (*models.RuleExecution, error) {
	rule, err := s.GetRuleByID(id)
	if err != nil {
		return nil, err
	}

	if !rule.Enabled {
		return nil, fmt.Errorf("rule %s is disabled", rule.Name)
	}

	execution := &models.RuleExecution{
		RuleID:      id,
		RuleName:    rule.Name,
		ExecutionID: uuid.New().String(),
		Status:      "RUNNING",
		StartTime:   time.Now(),
	}

	// Simulate rule execution
	// In a real implementation, this would evaluate conditions and execute actions
	time.Sleep(100 * time.Millisecond) // Simulate processing time

	execution.EndTime = time.Now()
	execution.Duration = execution.EndTime.Sub(execution.StartTime).Milliseconds()
	execution.Status = "COMPLETED"
	execution.Result = "Rule executed successfully"

	// Log execution (in a real implementation, store in database)
	return execution, nil
}

func (s *Service) GetRuleExecutions(ruleID string, limit, offset int) ([]models.RuleExecution, error) {
	// In a real implementation, this would query rule execution history from database
	// For now, return empty slice as no execution history is stored
	return []models.RuleExecution{}, nil
}
