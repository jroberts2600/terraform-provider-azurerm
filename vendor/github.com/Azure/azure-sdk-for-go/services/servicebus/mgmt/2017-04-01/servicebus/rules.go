package servicebus

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RulesClient is the azure Service Bus client
type RulesClient struct {
	BaseClient
}

// NewRulesClient creates an instance of the RulesClient client.
func NewRulesClient(subscriptionID string) RulesClient {
	return NewRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRulesClientWithBaseURI creates an instance of the RulesClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRulesClientWithBaseURI(baseURI string, subscriptionID string) RulesClient {
	return RulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new rule and updates an existing rule
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
// topicName - the topic name.
// subscriptionName - the subscription name.
// ruleName - the rule name.
// parameters - parameters supplied to create a rule.
func (client RulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string, parameters Rule) (result Rule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RulesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: topicName,
			Constraints: []validation.Constraint{{Target: "topicName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: subscriptionName,
			Constraints: []validation.Constraint{{Target: "subscriptionName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "subscriptionName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: ruleName,
			Constraints: []validation.Constraint{{Target: "ruleName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "ruleName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Ruleproperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Ruleproperties.SQLFilter", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Ruleproperties.SQLFilter.CompatibilityLevel", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Ruleproperties.SQLFilter.CompatibilityLevel", Name: validation.InclusiveMaximum, Rule: int64(20), Chain: nil},
							{Target: "parameters.Ruleproperties.SQLFilter.CompatibilityLevel", Name: validation.InclusiveMinimum, Rule: int64(20), Chain: nil},
						}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("servicebus.RulesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, namespaceName, topicName, subscriptionName, ruleName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.RulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.RulesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.RulesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string, parameters Rule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"subscriptionName":  autorest.Encode("path", subscriptionName),
		"topicName":         autorest.Encode("path", topicName),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RulesClient) CreateOrUpdateResponder(resp *http.Response) (result Rule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing rule.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
// topicName - the topic name.
// subscriptionName - the subscription name.
// ruleName - the rule name.
func (client RulesClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RulesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: topicName,
			Constraints: []validation.Constraint{{Target: "topicName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: subscriptionName,
			Constraints: []validation.Constraint{{Target: "subscriptionName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "subscriptionName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: ruleName,
			Constraints: []validation.Constraint{{Target: "ruleName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "ruleName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicebus.RulesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, namespaceName, topicName, subscriptionName, ruleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.RulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicebus.RulesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.RulesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"subscriptionName":  autorest.Encode("path", subscriptionName),
		"topicName":         autorest.Encode("path", topicName),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RulesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the description for the specified rule.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
// topicName - the topic name.
// subscriptionName - the subscription name.
// ruleName - the rule name.
func (client RulesClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string) (result Rule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RulesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: topicName,
			Constraints: []validation.Constraint{{Target: "topicName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: subscriptionName,
			Constraints: []validation.Constraint{{Target: "subscriptionName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "subscriptionName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: ruleName,
			Constraints: []validation.Constraint{{Target: "ruleName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "ruleName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicebus.RulesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, namespaceName, topicName, subscriptionName, ruleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.RulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.RulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.RulesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RulesClient) GetPreparer(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"subscriptionName":  autorest.Encode("path", subscriptionName),
		"topicName":         autorest.Encode("path", topicName),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RulesClient) GetResponder(resp *http.Response) (result Rule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscriptions list all the rules within given topic-subscription
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
// topicName - the topic name.
// subscriptionName - the subscription name.
// skip - skip is only used if a previous operation returned a partial result. If a previous response contains
// a nextLink element, the value of the nextLink element will include a skip parameter that specifies a
// starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N usageDetails.
func (client RulesClient) ListBySubscriptions(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, skip *int32, top *int32) (result RuleListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RulesClient.ListBySubscriptions")
		defer func() {
			sc := -1
			if result.rlr.Response.Response != nil {
				sc = result.rlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: topicName,
			Constraints: []validation.Constraint{{Target: "topicName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: subscriptionName,
			Constraints: []validation.Constraint{{Target: "subscriptionName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "subscriptionName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
				}}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("servicebus.RulesClient", "ListBySubscriptions", err.Error())
	}

	result.fn = client.listBySubscriptionsNextResults
	req, err := client.ListBySubscriptionsPreparer(ctx, resourceGroupName, namespaceName, topicName, subscriptionName, skip, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.RulesClient", "ListBySubscriptions", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionsSender(req)
	if err != nil {
		result.rlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.RulesClient", "ListBySubscriptions", resp, "Failure sending request")
		return
	}

	result.rlr, err = client.ListBySubscriptionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.RulesClient", "ListBySubscriptions", resp, "Failure responding to request")
		return
	}
	if result.rlr.hasNextLink() && result.rlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionsPreparer prepares the ListBySubscriptions request.
func (client RulesClient) ListBySubscriptionsPreparer(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, skip *int32, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"subscriptionName":  autorest.Encode("path", subscriptionName),
		"topicName":         autorest.Encode("path", topicName),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionsSender sends the ListBySubscriptions request. The method will close the
// http.Response Body if it receives an error.
func (client RulesClient) ListBySubscriptionsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionsResponder handles the response to the ListBySubscriptions request. The method always
// closes the http.Response Body.
func (client RulesClient) ListBySubscriptionsResponder(resp *http.Response) (result RuleListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionsNextResults retrieves the next set of results, if any.
func (client RulesClient) listBySubscriptionsNextResults(ctx context.Context, lastResults RuleListResult) (result RuleListResult, err error) {
	req, err := lastResults.ruleListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicebus.RulesClient", "listBySubscriptionsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicebus.RulesClient", "listBySubscriptionsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.RulesClient", "listBySubscriptionsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionsComplete enumerates all values, automatically crossing page boundaries as required.
func (client RulesClient) ListBySubscriptionsComplete(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, skip *int32, top *int32) (result RuleListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RulesClient.ListBySubscriptions")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscriptions(ctx, resourceGroupName, namespaceName, topicName, subscriptionName, skip, top)
	return
}
