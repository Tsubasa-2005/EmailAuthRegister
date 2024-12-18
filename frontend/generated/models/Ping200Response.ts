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
 * @interface Ping200Response
 */
export interface Ping200Response {
    /**
     * 
     * @type {string}
     * @memberof Ping200Response
     */
    message: string;
}

/**
 * Check if a given object implements the Ping200Response interface.
 */
export function instanceOfPing200Response(value: object): value is Ping200Response {
    if (!('message' in value) || value['message'] === undefined) return false;
    return true;
}

export function Ping200ResponseFromJSON(json: any): Ping200Response {
    return Ping200ResponseFromJSONTyped(json, false);
}

export function Ping200ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): Ping200Response {
    if (json == null) {
        return json;
    }
    return {
        
        'message': json['message'],
    };
}

export function Ping200ResponseToJSON(json: any): Ping200Response {
    return Ping200ResponseToJSONTyped(json, false);
}

export function Ping200ResponseToJSONTyped(value?: Ping200Response | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'message': value['message'],
    };
}

