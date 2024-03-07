package quivr_go_client

import "context"

type CreateBrainInput struct {
	Name               string                   `json:"name,omitempty"`
	Description        string                   `json:"description,omitempty"`
	Status             string                   `json:"status,omitempty"`
	Model              string                   `json:"model,omitempty"`
	Temperature        int                      `json:"temperature,omitempty"`
	MaxTokens          int                      `json:"max_tokens,omitempty"`
	PromptId           string                   `json:"prompt_id,omitempty"`
	BrainType          string                   `json:"brain_type,omitempty"`
	BrainDefinition    *BrainDefinitionInput    `json:"brain_definition,omitempty"`
	BrainSecretsValues *BrainSecretsValuesInput `json:"brain_secrets_values,omitempty"`
	ConnectedBrainsIds []string                 `json:"connected_brains_ids,omitempty"`
	Integration        *BrainIntegrationInput   `json:"integration,omitempty"`
}

type CreateBrainOutput struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Rights string `json:"rights"`
}

type BrainDefinitionInput struct {
	Method         string                      `json:"method,omitempty"`
	Url            string                      `json:"url,omitempty"`
	Params         *BrainDefinitionParamsInput `json:"params,omitempty"`
	SearchParams   *BrainDefinitionParamsInput `json:"search_params,omitempty"`
	Secrets        []interface{}               `json:"secrets,omitempty"`
	Raw            bool                        `json:"raw,omitempty"`
	JqInstructions string                      `json:"jq_instructions,omitempty"`
}

type BrainDefinitionParamsInput struct {
	Properties []string `json:"properties,omitempty"`
	Required   []string `json:"required,omitempty"`
}

type BrainSecretsValuesInput struct {
}

type BrainIntegrationInput struct {
	IntegrationId string                        `json:"integration_id,omitempty"`
	Settings      BrainIntegrationSettingsInput `json:"settings,omitempty"`
}

type BrainIntegrationSettingsInput struct {
}

func (c *Client) CreateBrain(ctx context.Context, input *CreateBrainInput) (*CreateBrainOutput, error) {
	return Do[CreateBrainOutput](ctx, c, POST, "/brains", input)
}
