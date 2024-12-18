/* tslint:disable */
/* eslint-disable */
/**
 * User Registration API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  CompleteUserRegistrationRequest,
  GetAllUsers200Response,
  Login400Response,
  LoginRequest,
  Ping200Response,
  SendEmailVerificationRequest,
  VerifyEmail200Response,
  VerifyEmailRequest,
} from '../models/index';
import {
    CompleteUserRegistrationRequestFromJSON,
    CompleteUserRegistrationRequestToJSON,
    GetAllUsers200ResponseFromJSON,
    GetAllUsers200ResponseToJSON,
    Login400ResponseFromJSON,
    Login400ResponseToJSON,
    LoginRequestFromJSON,
    LoginRequestToJSON,
    Ping200ResponseFromJSON,
    Ping200ResponseToJSON,
    SendEmailVerificationRequestFromJSON,
    SendEmailVerificationRequestToJSON,
    VerifyEmail200ResponseFromJSON,
    VerifyEmail200ResponseToJSON,
    VerifyEmailRequestFromJSON,
    VerifyEmailRequestToJSON,
} from '../models/index';

export interface CompleteUserRegistrationOperationRequest {
    completeUserRegistrationRequest: CompleteUserRegistrationRequest;
}

export interface GetAllUsersRequest {
    page: number;
}

export interface LoginOperationRequest {
    loginRequest: LoginRequest;
}

export interface SendEmailVerificationOperationRequest {
    sendEmailVerificationRequest: SendEmailVerificationRequest;
}

export interface VerifyEmailOperationRequest {
    verifyEmailRequest: VerifyEmailRequest;
}

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * Complete user registration
     */
    async completeUserRegistrationRaw(requestParameters: CompleteUserRegistrationOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['completeUserRegistrationRequest'] == null) {
            throw new runtime.RequiredError(
                'completeUserRegistrationRequest',
                'Required parameter "completeUserRegistrationRequest" was null or undefined when calling completeUserRegistration().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/complete-registration`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CompleteUserRegistrationRequestToJSON(requestParameters['completeUserRegistrationRequest']),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Complete user registration
     */
    async completeUserRegistration(requestParameters: CompleteUserRegistrationOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.completeUserRegistrationRaw(requestParameters, initOverrides);
    }

    /**
     * Get all users
     */
    async getAllUsersRaw(requestParameters: GetAllUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GetAllUsers200Response>> {
        if (requestParameters['page'] == null) {
            throw new runtime.RequiredError(
                'page',
                'Required parameter "page" was null or undefined when calling getAllUsers().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['page'] != null) {
            queryParameters['page'] = requestParameters['page'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("BearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/users`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GetAllUsers200ResponseFromJSON(jsonValue));
    }

    /**
     * Get all users
     */
    async getAllUsers(requestParameters: GetAllUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GetAllUsers200Response> {
        const response = await this.getAllUsersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Login
     */
    async loginRaw(requestParameters: LoginOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['loginRequest'] == null) {
            throw new runtime.RequiredError(
                'loginRequest',
                'Required parameter "loginRequest" was null or undefined when calling login().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/login`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: LoginRequestToJSON(requestParameters['loginRequest']),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Login
     */
    async login(requestParameters: LoginOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.loginRaw(requestParameters, initOverrides);
    }

    /**
     * Check if the server is running.
     * Ping the server
     */
    async pingRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Ping200Response>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/ping`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => Ping200ResponseFromJSON(jsonValue));
    }

    /**
     * Check if the server is running.
     * Ping the server
     */
    async ping(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Ping200Response> {
        const response = await this.pingRaw(initOverrides);
        return await response.value();
    }

    /**
     * Send email verification
     */
    async sendEmailVerificationRaw(requestParameters: SendEmailVerificationOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['sendEmailVerificationRequest'] == null) {
            throw new runtime.RequiredError(
                'sendEmailVerificationRequest',
                'Required parameter "sendEmailVerificationRequest" was null or undefined when calling sendEmailVerification().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/send-verification`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SendEmailVerificationRequestToJSON(requestParameters['sendEmailVerificationRequest']),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Send email verification
     */
    async sendEmailVerification(requestParameters: SendEmailVerificationOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.sendEmailVerificationRaw(requestParameters, initOverrides);
    }

    /**
     * Verify email
     */
    async verifyEmailRaw(requestParameters: VerifyEmailOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<VerifyEmail200Response>> {
        if (requestParameters['verifyEmailRequest'] == null) {
            throw new runtime.RequiredError(
                'verifyEmailRequest',
                'Required parameter "verifyEmailRequest" was null or undefined when calling verifyEmail().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/verify-email`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: VerifyEmailRequestToJSON(requestParameters['verifyEmailRequest']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => VerifyEmail200ResponseFromJSON(jsonValue));
    }

    /**
     * Verify email
     */
    async verifyEmail(requestParameters: VerifyEmailOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<VerifyEmail200Response> {
        const response = await this.verifyEmailRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
