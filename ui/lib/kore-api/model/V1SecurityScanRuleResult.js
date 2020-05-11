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

/**
 * The V1SecurityScanRuleResult model module.
 * @module model/V1SecurityScanRuleResult
 * @version 0.0.1
 */
class V1SecurityScanRuleResult {
    /**
     * Constructs a new <code>V1SecurityScanRuleResult</code>.
     * @alias module:model/V1SecurityScanRuleResult
     */
    constructor() { 
        
        V1SecurityScanRuleResult.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1SecurityScanRuleResult</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1SecurityScanRuleResult} obj Optional instance to populate.
     * @return {module:model/V1SecurityScanRuleResult} The populated <code>V1SecurityScanRuleResult</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1SecurityScanRuleResult();

            if (data.hasOwnProperty('checkedAt')) {
                obj['checkedAt'] = ApiClient.convertToType(data['checkedAt'], 'String');
            }
            if (data.hasOwnProperty('message')) {
                obj['message'] = ApiClient.convertToType(data['message'], 'String');
            }
            if (data.hasOwnProperty('ruleCode')) {
                obj['ruleCode'] = ApiClient.convertToType(data['ruleCode'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {String}
     */
    getCheckedAt() {
        return this.checkedAt;
    }

    /**
     * @param {String} checkedAt
     */
    setCheckedAt(checkedAt) {
        this['checkedAt'] = checkedAt;
    }
/**
     * @return {String}
     */
    getMessage() {
        return this.message;
    }

    /**
     * @param {String} message
     */
    setMessage(message) {
        this['message'] = message;
    }
/**
     * @return {String}
     */
    getRuleCode() {
        return this.ruleCode;
    }

    /**
     * @param {String} ruleCode
     */
    setRuleCode(ruleCode) {
        this['ruleCode'] = ruleCode;
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

}

/**
 * @member {String} checkedAt
 */
V1SecurityScanRuleResult.prototype['checkedAt'] = undefined;

/**
 * @member {String} message
 */
V1SecurityScanRuleResult.prototype['message'] = undefined;

/**
 * @member {String} ruleCode
 */
V1SecurityScanRuleResult.prototype['ruleCode'] = undefined;

/**
 * @member {String} status
 */
V1SecurityScanRuleResult.prototype['status'] = undefined;






export default V1SecurityScanRuleResult;
