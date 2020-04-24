/**
 * Appvia Kore API
 * Kore API provides the frontend API for the Appvia Kore (kore.appvia.io)
 *
 * The version of the OpenAPI document: 0.0.1
 * Contact: info@appvia.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1ServicePlanSpec from './V1ServicePlanSpec';

/**
 * The ApiserverTeamServicePlan model module.
 * @module model/ApiserverTeamServicePlan
 * @version 0.0.1
 */
class ApiserverTeamServicePlan {
    /**
     * Constructs a new <code>ApiserverTeamServicePlan</code>.
     * @alias module:model/ApiserverTeamServicePlan
     * @param parameterEditable {Object.<String, Boolean>} 
     * @param schema {String} 
     */
    constructor(parameterEditable, schema) { 
        
        ApiserverTeamServicePlan.initialize(this, parameterEditable, schema);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, parameterEditable, schema) { 
        obj['parameterEditable'] = parameterEditable;
        obj['schema'] = schema;
    }

    /**
     * Constructs a <code>ApiserverTeamServicePlan</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiserverTeamServicePlan} obj Optional instance to populate.
     * @return {module:model/ApiserverTeamServicePlan} The populated <code>ApiserverTeamServicePlan</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiserverTeamServicePlan();

            if (data.hasOwnProperty('parameterEditable')) {
                obj['parameterEditable'] = ApiClient.convertToType(data['parameterEditable'], {'String': 'Boolean'});
            }
            if (data.hasOwnProperty('schema')) {
                obj['schema'] = ApiClient.convertToType(data['schema'], 'String');
            }
            if (data.hasOwnProperty('servicePlan')) {
                obj['servicePlan'] = V1ServicePlanSpec.constructFromObject(data['servicePlan']);
            }
        }
        return obj;
    }

/**
     * @return {Object.<String, Boolean>}
     */
    getParameterEditable() {
        return this.parameterEditable;
    }

    /**
     * @param {Object.<String, Boolean>} parameterEditable
     */
    setParameterEditable(parameterEditable) {
        this['parameterEditable'] = parameterEditable;
    }
/**
     * @return {String}
     */
    getSchema() {
        return this.schema;
    }

    /**
     * @param {String} schema
     */
    setSchema(schema) {
        this['schema'] = schema;
    }
/**
     * @return {module:model/V1ServicePlanSpec}
     */
    getServicePlan() {
        return this.servicePlan;
    }

    /**
     * @param {module:model/V1ServicePlanSpec} servicePlan
     */
    setServicePlan(servicePlan) {
        this['servicePlan'] = servicePlan;
    }

}

/**
 * @member {Object.<String, Boolean>} parameterEditable
 */
ApiserverTeamServicePlan.prototype['parameterEditable'] = undefined;

/**
 * @member {String} schema
 */
ApiserverTeamServicePlan.prototype['schema'] = undefined;

/**
 * @member {module:model/V1ServicePlanSpec} servicePlan
 */
ApiserverTeamServicePlan.prototype['servicePlan'] = undefined;






export default ApiserverTeamServicePlan;

