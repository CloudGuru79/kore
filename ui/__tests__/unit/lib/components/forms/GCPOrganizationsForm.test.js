import { mount } from 'enzyme'

import GCPOrganizationForm from '../../../../../lib/components/forms/GCPOrganizationForm'
import ApiTestHelpers from '../../../../api-test-helpers'

describe('GCPOrganizationForm', () => {
  let props
  let form
  let apiScope
  const secret = {
    metadata: { name: 'gcp' },
    spec: { type: 'gcp-org' }
  }
  const gcpOrganization = {
    metadata: { name: 'gcp' },
    spec: { parentID: 'org-id', billingAccount: 'billing@example.com', account: 'org-cred' }
  }
  const allocation = {
    metadata: { name: 'gcp' },
    spec: { resource: { kind: 'Organization' } }
  }

  beforeEach(() => {
    // In case any tests leak to the API, mock out the API at this stage:
    apiScope = (ApiTestHelpers.getScope())

    props = {
      form: {
        isFieldTouched: () => {},
        getFieldDecorator: jest.fn(() => c => c),
        getFieldsError: () => () => {},
        getFieldError: () => {},
        getFieldValue: () => {},
        validateFields: jest.fn()
      },
      team: 'abc',
      allTeams: { items: [] },
      handleSubmit: jest.fn()
    }
    mount(<GCPOrganizationForm wrappedComponentRef={component => form = component} {...props} />)
  })

  afterEach(() => {
    // This will check that no calls were made against the API, unless the test registered them:
    apiScope.done()
  })

  describe('#generateSecretResource', () => {
    it('returns a configured Secret object', () => {
      const secret = form.generateSecretResource({ name: 'gcp', parentID: 'appvia', account: 'gcp-cred' })
      expect(secret).toBeDefined()
    })
  })

  describe('#generateGCPOrganizationResource', () => {
    it('returns a configured Organization object when given valid values', () => {
      const gcpOrg = form.generateGCPOrganizationResource({ name: 'gcp', parentID: 'org-id', billingAccount: 'billing@example.com' })
      expect(gcpOrg).toBeDefined()
    })
  })

  describe('#getResource', () => {
    beforeEach(() => {
      apiScope
        .get(`${ApiTestHelpers.basePath}/teams/abc/organizations/gcp`).reply(200, gcpOrganization)
        .get(`${ApiTestHelpers.basePath}/teams/abc/allocations/gcp`).reply(200, allocation)
    })

    it('returns Organization and allocation from API', async () => {
      const result = await form.getResource('gcp')
      const expected = { ...gcpOrganization, allocation }
      expect(result).toEqual(expected)
      apiScope.done()
    })
  })

  describe('#putResource', () => {
    beforeEach(() => {
      apiScope
        .put(`${ApiTestHelpers.basePath}/teams/abc/secrets/gcp`).reply(200, secret)
        .put(`${ApiTestHelpers.basePath}/teams/abc/organizations/gcp`).reply(200, gcpOrganization)
        .put(`${ApiTestHelpers.basePath}/teams/abc/allocations/gcp`).reply(200, allocation)
    })

    it('creates/updates and returns Organization and allocation from API', async () => {
      const result = await form.putResource({ name: 'gcp', project: 'project-id', account: 'gke-service-account-cred' })
      const expected = { ...gcpOrganization, allocation }
      expect(result).toEqual(expected)
      apiScope.done()
    })
  })
})