// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package contactcenterinsights aliases all exported identifiers in package
// "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb".
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package contactcenterinsights

import (
	src "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
const (
	AnswerFeedback_CORRECTNESS_LEVEL_UNSPECIFIED                  = src.AnswerFeedback_CORRECTNESS_LEVEL_UNSPECIFIED
	AnswerFeedback_FULLY_CORRECT                                  = src.AnswerFeedback_FULLY_CORRECT
	AnswerFeedback_NOT_CORRECT                                    = src.AnswerFeedback_NOT_CORRECT
	AnswerFeedback_PARTIALLY_CORRECT                              = src.AnswerFeedback_PARTIALLY_CORRECT
	ConversationParticipant_ANY_AGENT                             = src.ConversationParticipant_ANY_AGENT
	ConversationParticipant_AUTOMATED_AGENT                       = src.ConversationParticipant_AUTOMATED_AGENT
	ConversationParticipant_END_USER                              = src.ConversationParticipant_END_USER
	ConversationParticipant_HUMAN_AGENT                           = src.ConversationParticipant_HUMAN_AGENT
	ConversationParticipant_ROLE_UNSPECIFIED                      = src.ConversationParticipant_ROLE_UNSPECIFIED
	ConversationView_BASIC                                        = src.ConversationView_BASIC
	ConversationView_CONVERSATION_VIEW_UNSPECIFIED                = src.ConversationView_CONVERSATION_VIEW_UNSPECIFIED
	ConversationView_FULL                                         = src.ConversationView_FULL
	Conversation_CHAT                                             = src.Conversation_CHAT
	Conversation_MEDIUM_UNSPECIFIED                               = src.Conversation_MEDIUM_UNSPECIFIED
	Conversation_PHONE_CALL                                       = src.Conversation_PHONE_CALL
	EntityMentionData_COMMON                                      = src.EntityMentionData_COMMON
	EntityMentionData_MENTION_TYPE_UNSPECIFIED                    = src.EntityMentionData_MENTION_TYPE_UNSPECIFIED
	EntityMentionData_PROPER                                      = src.EntityMentionData_PROPER
	Entity_ADDRESS                                                = src.Entity_ADDRESS
	Entity_CONSUMER_GOOD                                          = src.Entity_CONSUMER_GOOD
	Entity_DATE                                                   = src.Entity_DATE
	Entity_EVENT                                                  = src.Entity_EVENT
	Entity_LOCATION                                               = src.Entity_LOCATION
	Entity_NUMBER                                                 = src.Entity_NUMBER
	Entity_ORGANIZATION                                           = src.Entity_ORGANIZATION
	Entity_OTHER                                                  = src.Entity_OTHER
	Entity_PERSON                                                 = src.Entity_PERSON
	Entity_PHONE_NUMBER                                           = src.Entity_PHONE_NUMBER
	Entity_PRICE                                                  = src.Entity_PRICE
	Entity_TYPE_UNSPECIFIED                                       = src.Entity_TYPE_UNSPECIFIED
	Entity_WORK_OF_ART                                            = src.Entity_WORK_OF_ART
	ExportInsightsDataRequest_WRITE_APPEND                        = src.ExportInsightsDataRequest_WRITE_APPEND
	ExportInsightsDataRequest_WRITE_DISPOSITION_UNSPECIFIED       = src.ExportInsightsDataRequest_WRITE_DISPOSITION_UNSPECIFIED
	ExportInsightsDataRequest_WRITE_TRUNCATE                      = src.ExportInsightsDataRequest_WRITE_TRUNCATE
	IssueModel_DELETING                                           = src.IssueModel_DELETING
	IssueModel_DEPLOYED                                           = src.IssueModel_DEPLOYED
	IssueModel_DEPLOYING                                          = src.IssueModel_DEPLOYING
	IssueModel_STATE_UNSPECIFIED                                  = src.IssueModel_STATE_UNSPECIFIED
	IssueModel_UNDEPLOYED                                         = src.IssueModel_UNDEPLOYED
	IssueModel_UNDEPLOYING                                        = src.IssueModel_UNDEPLOYING
	PhraseMatchRuleGroup_ALL_OF                                   = src.PhraseMatchRuleGroup_ALL_OF
	PhraseMatchRuleGroup_ANY_OF                                   = src.PhraseMatchRuleGroup_ANY_OF
	PhraseMatchRuleGroup_PHRASE_MATCH_RULE_GROUP_TYPE_UNSPECIFIED = src.PhraseMatchRuleGroup_PHRASE_MATCH_RULE_GROUP_TYPE_UNSPECIFIED
	PhraseMatcher_ALL_OF                                          = src.PhraseMatcher_ALL_OF
	PhraseMatcher_ANY_OF                                          = src.PhraseMatcher_ANY_OF
	PhraseMatcher_PHRASE_MATCHER_TYPE_UNSPECIFIED                 = src.PhraseMatcher_PHRASE_MATCHER_TYPE_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
var (
	AnswerFeedback_CorrectnessLevel_name                                     = src.AnswerFeedback_CorrectnessLevel_name
	AnswerFeedback_CorrectnessLevel_value                                    = src.AnswerFeedback_CorrectnessLevel_value
	ConversationParticipant_Role_name                                        = src.ConversationParticipant_Role_name
	ConversationParticipant_Role_value                                       = src.ConversationParticipant_Role_value
	ConversationView_name                                                    = src.ConversationView_name
	ConversationView_value                                                   = src.ConversationView_value
	Conversation_Medium_name                                                 = src.Conversation_Medium_name
	Conversation_Medium_value                                                = src.Conversation_Medium_value
	EntityMentionData_MentionType_name                                       = src.EntityMentionData_MentionType_name
	EntityMentionData_MentionType_value                                      = src.EntityMentionData_MentionType_value
	Entity_Type_name                                                         = src.Entity_Type_name
	Entity_Type_value                                                        = src.Entity_Type_value
	ExportInsightsDataRequest_WriteDisposition_name                          = src.ExportInsightsDataRequest_WriteDisposition_name
	ExportInsightsDataRequest_WriteDisposition_value                         = src.ExportInsightsDataRequest_WriteDisposition_value
	File_google_cloud_contactcenterinsights_v1_contact_center_insights_proto = src.File_google_cloud_contactcenterinsights_v1_contact_center_insights_proto
	File_google_cloud_contactcenterinsights_v1_resources_proto               = src.File_google_cloud_contactcenterinsights_v1_resources_proto
	IssueModel_State_name                                                    = src.IssueModel_State_name
	IssueModel_State_value                                                   = src.IssueModel_State_value
	PhraseMatchRuleGroup_PhraseMatchRuleGroupType_name                       = src.PhraseMatchRuleGroup_PhraseMatchRuleGroupType_name
	PhraseMatchRuleGroup_PhraseMatchRuleGroupType_value                      = src.PhraseMatchRuleGroup_PhraseMatchRuleGroupType_value
	PhraseMatcher_PhraseMatcherType_name                                     = src.PhraseMatcher_PhraseMatcherType_name
	PhraseMatcher_PhraseMatcherType_value                                    = src.PhraseMatcher_PhraseMatcherType_value
)

// The analysis resource.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Analysis = src.Analysis

// The result of an analysis.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type AnalysisResult = src.AnalysisResult

// Call-specific metadata created during analysis.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type AnalysisResult_CallAnalysisMetadata = src.AnalysisResult_CallAnalysisMetadata
type AnalysisResult_CallAnalysisMetadata_ = src.AnalysisResult_CallAnalysisMetadata_

// A point in a conversation that marks the start or the end of an annotation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type AnnotationBoundary = src.AnnotationBoundary
type AnnotationBoundary_WordIndex = src.AnnotationBoundary_WordIndex

// The feedback that the customer has about a certain answer in the
// conversation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type AnswerFeedback = src.AnswerFeedback

// The correctness level of an answer.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type AnswerFeedback_CorrectnessLevel = src.AnswerFeedback_CorrectnessLevel

// Agent Assist Article Suggestion data.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ArticleSuggestionData = src.ArticleSuggestionData

// Request to get statistics of an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CalculateIssueModelStatsRequest = src.CalculateIssueModelStatsRequest

// Response of querying an issue model's statistics.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CalculateIssueModelStatsResponse = src.CalculateIssueModelStatsResponse

// The request for calculating conversation statistics.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CalculateStatsRequest = src.CalculateStatsRequest

// The response for calculating conversation statistics.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CalculateStatsResponse = src.CalculateStatsResponse

// A time series representing conversations over time.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CalculateStatsResponse_TimeSeries = src.CalculateStatsResponse_TimeSeries

// A single interval in a time series.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CalculateStatsResponse_TimeSeries_Interval = src.CalculateStatsResponse_TimeSeries_Interval

// A piece of metadata that applies to a window of a call.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CallAnnotation = src.CallAnnotation
type CallAnnotation_EntityMentionData = src.CallAnnotation_EntityMentionData
type CallAnnotation_HoldData = src.CallAnnotation_HoldData
type CallAnnotation_IntentMatchData = src.CallAnnotation_IntentMatchData
type CallAnnotation_InterruptionData = src.CallAnnotation_InterruptionData
type CallAnnotation_PhraseMatchData = src.CallAnnotation_PhraseMatchData
type CallAnnotation_SentimentData = src.CallAnnotation_SentimentData
type CallAnnotation_SilenceData = src.CallAnnotation_SilenceData

// ContactCenterInsightsClient is the client API for ContactCenterInsights
// service. For semantics around ctx use and closing/ending streaming RPCs,
// please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ContactCenterInsightsClient = src.ContactCenterInsightsClient

// ContactCenterInsightsServer is the server API for ContactCenterInsights
// service.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ContactCenterInsightsServer = src.ContactCenterInsightsServer

// The conversation resource.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Conversation = src.Conversation

// The conversation source, which is a combination of transcript and audio.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ConversationDataSource = src.ConversationDataSource
type ConversationDataSource_DialogflowSource = src.ConversationDataSource_DialogflowSource
type ConversationDataSource_GcsSource = src.ConversationDataSource_GcsSource

// One channel of conversation-level sentiment data.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ConversationLevelSentiment = src.ConversationLevelSentiment

// The call participant speaking for a given utterance.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ConversationParticipant = src.ConversationParticipant
type ConversationParticipant_DialogflowParticipantName = src.ConversationParticipant_DialogflowParticipantName

// The role of the participant.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ConversationParticipant_Role = src.ConversationParticipant_Role
type ConversationParticipant_UserId = src.ConversationParticipant_UserId

// Represents the options for viewing a conversation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ConversationView = src.ConversationView

// Call-specific metadata.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Conversation_CallMetadata = src.Conversation_CallMetadata
type Conversation_CallMetadata_ = src.Conversation_CallMetadata_
type Conversation_ExpireTime = src.Conversation_ExpireTime

// Possible media for the conversation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Conversation_Medium = src.Conversation_Medium

// A message representing the transcript of a conversation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Conversation_Transcript = src.Conversation_Transcript

// A segment of a full transcript.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Conversation_Transcript_TranscriptSegment = src.Conversation_Transcript_TranscriptSegment

// Metadata from Dialogflow relating to the current transcript segment.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Conversation_Transcript_TranscriptSegment_DialogflowSegmentMetadata = src.Conversation_Transcript_TranscriptSegment_DialogflowSegmentMetadata

// Word-level info for words in a transcript.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Conversation_Transcript_TranscriptSegment_WordInfo = src.Conversation_Transcript_TranscriptSegment_WordInfo
type Conversation_Ttl = src.Conversation_Ttl

// Metadata for a create analysis operation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CreateAnalysisOperationMetadata = src.CreateAnalysisOperationMetadata

// The request to create an analysis.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CreateAnalysisRequest = src.CreateAnalysisRequest

// Request to create a conversation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CreateConversationRequest = src.CreateConversationRequest

// Metadata for creating an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CreateIssueModelMetadata = src.CreateIssueModelMetadata

// The request to create an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CreateIssueModelRequest = src.CreateIssueModelRequest

// Request to create a phrase matcher.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CreatePhraseMatcherRequest = src.CreatePhraseMatcherRequest

// The request to create a view.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type CreateViewRequest = src.CreateViewRequest

// The request to delete an analysis.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type DeleteAnalysisRequest = src.DeleteAnalysisRequest

// The request to delete a conversation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type DeleteConversationRequest = src.DeleteConversationRequest

// Metadata for deleting an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type DeleteIssueModelMetadata = src.DeleteIssueModelMetadata

// The request to delete an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type DeleteIssueModelRequest = src.DeleteIssueModelRequest

// The request to delete a phrase matcher.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type DeletePhraseMatcherRequest = src.DeletePhraseMatcherRequest

// The request to delete a view.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type DeleteViewRequest = src.DeleteViewRequest

// Metadata for deploying an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type DeployIssueModelMetadata = src.DeployIssueModelMetadata

// The request to deploy an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type DeployIssueModelRequest = src.DeployIssueModelRequest

// The response to deploy an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type DeployIssueModelResponse = src.DeployIssueModelResponse

// The data for a Dialogflow intent. Represents a detected intent in the
// conversation, e.g. MAKES_PROMISE.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type DialogflowIntent = src.DialogflowIntent

// Dialogflow interaction data.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type DialogflowInteractionData = src.DialogflowInteractionData

// A Dialogflow source of conversation data.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type DialogflowSource = src.DialogflowSource

// The data for an entity annotation. Represents a phrase in the conversation
// that is a known entity, such as a person, an organization, or location.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Entity = src.Entity

// The data for an entity mention annotation. This represents a mention of an
// `Entity` in the conversation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type EntityMentionData = src.EntityMentionData

// The supported types of mentions.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type EntityMentionData_MentionType = src.EntityMentionData_MentionType

// The type of the entity. For most entity types, the associated metadata is a
// Wikipedia URL (`wikipedia_url`) and Knowledge Graph MID (`mid`). The table
// below lists the associated fields for entities that have different metadata.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Entity_Type = src.Entity_Type

// Exact match configuration.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ExactMatchConfig = src.ExactMatchConfig

// Metadata for an export insights operation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ExportInsightsDataMetadata = src.ExportInsightsDataMetadata

// The request to export insights.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ExportInsightsDataRequest = src.ExportInsightsDataRequest

// A BigQuery Table Reference.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ExportInsightsDataRequest_BigQueryDestination = src.ExportInsightsDataRequest_BigQueryDestination
type ExportInsightsDataRequest_BigQueryDestination_ = src.ExportInsightsDataRequest_BigQueryDestination_

// Specifies the action that occurs if the destination table already exists.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ExportInsightsDataRequest_WriteDisposition = src.ExportInsightsDataRequest_WriteDisposition

// Response for an export insights operation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ExportInsightsDataResponse = src.ExportInsightsDataResponse

// Agent Assist frequently-asked-question answer data.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type FaqAnswerData = src.FaqAnswerData

// A Cloud Storage source of conversation data.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type GcsSource = src.GcsSource

// The request to get an analysis.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type GetAnalysisRequest = src.GetAnalysisRequest

// The request to get a conversation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type GetConversationRequest = src.GetConversationRequest

// The request to get an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type GetIssueModelRequest = src.GetIssueModelRequest

// The request to get an issue.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type GetIssueRequest = src.GetIssueRequest

// The request to get a a phrase matcher.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type GetPhraseMatcherRequest = src.GetPhraseMatcherRequest

// The request to get project-level settings.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type GetSettingsRequest = src.GetSettingsRequest

// The request to get a view.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type GetViewRequest = src.GetViewRequest

// The data for a hold annotation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type HoldData = src.HoldData

// The data for an intent. Represents a detected intent in the conversation,
// for example MAKES_PROMISE.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Intent = src.Intent

// The data for an intent match. Represents an intent match for a text segment
// in the conversation. A text segment can be part of a sentence, a complete
// sentence, or an utterance with multiple sentences.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type IntentMatchData = src.IntentMatchData

// The data for an interruption annotation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type InterruptionData = src.InterruptionData

// The issue resource.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Issue = src.Issue

// Information about the issue.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type IssueAssignment = src.IssueAssignment

// The issue model resource.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type IssueModel = src.IssueModel

// Aggregated statistics about an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type IssueModelLabelStats = src.IssueModelLabelStats

// Aggregated statistics about an issue.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type IssueModelLabelStats_IssueStats = src.IssueModelLabelStats_IssueStats

// Issue Modeling result on a conversation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type IssueModelResult = src.IssueModelResult

// Configs for the input data used to create the issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type IssueModel_InputDataConfig = src.IssueModel_InputDataConfig

// State of the model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type IssueModel_State = src.IssueModel_State

// The request to list analyses.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ListAnalysesRequest = src.ListAnalysesRequest

// The response to list analyses.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ListAnalysesResponse = src.ListAnalysesResponse

// Request to list conversations.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ListConversationsRequest = src.ListConversationsRequest

// The response of listing conversations.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ListConversationsResponse = src.ListConversationsResponse

// Request to list issue models.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ListIssueModelsRequest = src.ListIssueModelsRequest

// The response of listing issue models.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ListIssueModelsResponse = src.ListIssueModelsResponse

// Request to list issues.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ListIssuesRequest = src.ListIssuesRequest

// The response of listing issues.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ListIssuesResponse = src.ListIssuesResponse

// Request to list phrase matchers.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ListPhraseMatchersRequest = src.ListPhraseMatchersRequest

// The response of listing phrase matchers.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ListPhraseMatchersResponse = src.ListPhraseMatchersResponse

// The request to list views.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ListViewsRequest = src.ListViewsRequest

// The response of listing views.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type ListViewsResponse = src.ListViewsResponse

// The data for a matched phrase matcher. Represents information identifying a
// phrase matcher for a given match.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type PhraseMatchData = src.PhraseMatchData

// The data for a phrase match rule.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type PhraseMatchRule = src.PhraseMatchRule

// Configuration information of a phrase match rule.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type PhraseMatchRuleConfig = src.PhraseMatchRuleConfig
type PhraseMatchRuleConfig_ExactMatchConfig = src.PhraseMatchRuleConfig_ExactMatchConfig

// A message representing a rule in the phrase matcher.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type PhraseMatchRuleGroup = src.PhraseMatchRuleGroup

// Specifies how to combine each phrase match rule for whether there is a
// match.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type PhraseMatchRuleGroup_PhraseMatchRuleGroupType = src.PhraseMatchRuleGroup_PhraseMatchRuleGroupType

// The phrase matcher resource.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type PhraseMatcher = src.PhraseMatcher

// Specifies how to combine each phrase match rule group to determine whether
// there is a match.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type PhraseMatcher_PhraseMatcherType = src.PhraseMatcher_PhraseMatcherType

// An annotation that was generated during the customer and agent interaction.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type RuntimeAnnotation = src.RuntimeAnnotation
type RuntimeAnnotation_ArticleSuggestion = src.RuntimeAnnotation_ArticleSuggestion
type RuntimeAnnotation_DialogflowInteraction = src.RuntimeAnnotation_DialogflowInteraction
type RuntimeAnnotation_FaqAnswer = src.RuntimeAnnotation_FaqAnswer
type RuntimeAnnotation_SmartComposeSuggestion = src.RuntimeAnnotation_SmartComposeSuggestion
type RuntimeAnnotation_SmartReply = src.RuntimeAnnotation_SmartReply

// The data for a sentiment annotation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type SentimentData = src.SentimentData

// The settings resource.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Settings = src.Settings

// Default configuration when creating Analyses in Insights.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type Settings_AnalysisConfig = src.Settings_AnalysisConfig

// The data for a silence annotation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type SilenceData = src.SilenceData

// Agent Assist Smart Compose suggestion data.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type SmartComposeSuggestionData = src.SmartComposeSuggestionData

// Agent Assist Smart Reply data.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type SmartReplyData = src.SmartReplyData

// Metadata for undeploying an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type UndeployIssueModelMetadata = src.UndeployIssueModelMetadata

// The request to undeploy an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type UndeployIssueModelRequest = src.UndeployIssueModelRequest

// The response to undeploy an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type UndeployIssueModelResponse = src.UndeployIssueModelResponse

// UnimplementedContactCenterInsightsServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type UnimplementedContactCenterInsightsServer = src.UnimplementedContactCenterInsightsServer

// The request to update a conversation.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type UpdateConversationRequest = src.UpdateConversationRequest

// The request to update an issue model.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type UpdateIssueModelRequest = src.UpdateIssueModelRequest

// The request to update an issue.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type UpdateIssueRequest = src.UpdateIssueRequest

// The request to update a phrase matcher.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type UpdatePhraseMatcherRequest = src.UpdatePhraseMatcherRequest

// The request to update project-level settings.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type UpdateSettingsRequest = src.UpdateSettingsRequest

// The request to update a view.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type UpdateViewRequest = src.UpdateViewRequest

// The View resource.
//
// Deprecated: Please use types in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
type View = src.View

// Deprecated: Please use funcs in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
func NewContactCenterInsightsClient(cc grpc.ClientConnInterface) ContactCenterInsightsClient {
	return src.NewContactCenterInsightsClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb
func RegisterContactCenterInsightsServer(s *grpc.Server, srv ContactCenterInsightsServer) {
	src.RegisterContactCenterInsightsServer(s, srv)
}