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
import V1Condition from './V1Condition';

/**
 * The V1alpha1EKSCredentialsStatus model module.
 * @module model/V1alpha1EKSCredentialsStatus
 * @version 0.0.1
 */
class V1alpha1EKSCredentialsStatus {
    /**
     * Constructs a new <code>V1alpha1EKSCredentialsStatus</code>.
     * @alias module:model/V1alpha1EKSCredentialsStatus
     */
    constructor() { 
        
        V1alpha1EKSCredentialsStatus.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1alpha1EKSCredentialsStatus</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1alpha1EKSCredentialsStatus} obj Optional instance to populate.
     * @return {module:model/V1alpha1EKSCredentialsStatus} The populated <code>V1alpha1EKSCredentialsStatus</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1alpha1EKSCredentialsStatus();

            if (data.hasOwnProperty('conditions')) {
                obj['conditions'] = ApiClient.convertToType(data['conditions'], [V1Condition]);
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('verified')) {
                obj['verified'] = ApiClient.convertToType(data['verified'], 'Boolean');
            }
        }
        return obj;
    }

/**
     * @return {Array.<module:model/V1Condition>}
     */
    getConditions() {
        return this.conditions;
    }

    /**
     * @param {Array.<module:model/V1Condition>} conditions
     */
    setConditions(conditions) {
        this['conditions'] = conditions;
    }
/**
     * @return {String}
     */
    getStatus() {
        return this.status;
    }

    /**
     * @param {String} status
     */
    setStatus(status) {
        this['status'] = status;
    }
/**
     * @return {Boolean}
     */
    getVerified() {
        return this.verified;
    }

    /**
     * @param {Boolean} verified
     */
    setVerified(verified) {
        this['verified'] = verified;
    }

}

/**
 * @member {Array.<module:model/V1Condition>} conditions
 */
V1alpha1EKSCredentialsStatus.prototype['conditions'] = undefined;

/**
 * @member {String} status
 */
V1alpha1EKSCredentialsStatus.prototype['status'] = undefined;

/**
 * @member {Boolean} verified
 */
V1alpha1EKSCredentialsStatus.prototype['verified'] = undefined;






export default V1alpha1EKSCredentialsStatus;
