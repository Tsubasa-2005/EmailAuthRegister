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
import type { User } from './User';
import {
    UserFromJSON,
    UserFromJSONTyped,
    UserToJSON,
    UserToJSONTyped,
} from './User';

/**
 * 
 * @export
 * @interface GetAllUsers200Response
 */
export interface GetAllUsers200Response {
    /**
     * 
     * @type {Array<User>}
     * @memberof GetAllUsers200Response
     */
    users: Array<User>;
    /**
     * Total number of pages
     * @type {number}
     * @memberof GetAllUsers200Response
     */
    totalPage: number;
}

/**
 * Check if a given object implements the GetAllUsers200Response interface.
 */
export function instanceOfGetAllUsers200Response(value: object): value is GetAllUsers200Response {
    if (!('users' in value) || value['users'] === undefined) return false;
    if (!('totalPage' in value) || value['totalPage'] === undefined) return false;
    return true;
}

export function GetAllUsers200ResponseFromJSON(json: any): GetAllUsers200Response {
    return GetAllUsers200ResponseFromJSONTyped(json, false);
}

export function GetAllUsers200ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetAllUsers200Response {
    if (json == null) {
        return json;
    }
    return {
        
        'users': ((json['users'] as Array<any>).map(UserFromJSON)),
        'totalPage': json['totalPage'],
    };
}

export function GetAllUsers200ResponseToJSON(json: any): GetAllUsers200Response {
    return GetAllUsers200ResponseToJSONTyped(json, false);
}

export function GetAllUsers200ResponseToJSONTyped(value?: GetAllUsers200Response | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'users': ((value['users'] as Array<any>).map(UserToJSON)),
        'totalPage': value['totalPage'],
    };
}

