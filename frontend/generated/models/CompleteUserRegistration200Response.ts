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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface CompleteUserRegistration200Response
 */
export interface CompleteUserRegistration200Response {
    /**
     * 
     * @type {string}
     * @memberof CompleteUserRegistration200Response
     */
    token: string;
}

/**
 * Check if a given object implements the CompleteUserRegistration200Response interface.
 */
export function instanceOfCompleteUserRegistration200Response(value: object): value is CompleteUserRegistration200Response {
    if (!('token' in value) || value['token'] === undefined) return false;
    return true;
}

export function CompleteUserRegistration200ResponseFromJSON(json: any): CompleteUserRegistration200Response {
    return CompleteUserRegistration200ResponseFromJSONTyped(json, false);
}

export function CompleteUserRegistration200ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CompleteUserRegistration200Response {
    if (json == null) {
        return json;
    }
    return {
        
        'token': json['token'],
    };
}

export function CompleteUserRegistration200ResponseToJSON(json: any): CompleteUserRegistration200Response {
    return CompleteUserRegistration200ResponseToJSONTyped(json, false);
}

export function CompleteUserRegistration200ResponseToJSONTyped(value?: CompleteUserRegistration200Response | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'token': value['token'],
    };
}
