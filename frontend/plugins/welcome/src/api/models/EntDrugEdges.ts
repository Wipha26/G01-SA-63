/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDispense,
    EntDispenseFromJSON,
    EntDispenseFromJSONTyped,
    EntDispenseToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDrugEdges
 */
export interface EntDrugEdges {
    /**
     * Dispenses holds the value of the dispenses edge.
     * @type {Array<EntDispense>}
     * @memberof EntDrugEdges
     */
    dispenses?: Array<EntDispense>;
}

export function EntDrugEdgesFromJSON(json: any): EntDrugEdges {
    return EntDrugEdgesFromJSONTyped(json, false);
}

export function EntDrugEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDrugEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dispenses': !exists(json, 'dispenses') ? undefined : ((json['dispenses'] as Array<any>).map(EntDispenseFromJSON)),
    };
}

export function EntDrugEdgesToJSON(value?: EntDrugEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'dispenses': value.dispenses === undefined ? undefined : ((value.dispenses as Array<any>).map(EntDispenseToJSON)),
    };
}

