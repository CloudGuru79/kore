import axios from 'axios'
import App from '../../pages/_app'
import OrgService from '../../server/services/org'
import apiRequest from '../../lib/utils/api-request'
import redirect from '../../lib/utils/redirect'

jest.mock('../../server/services/org')
jest.mock('../../lib/utils/api-request')
jest.mock('../../lib/utils/redirect')
jest.mock('axios')

describe('App', () => {
  beforeEach(() => {
    OrgService.mockClear()
    apiRequest.mockClear()
    redirect.mockClear()
    axios.mockClear()
  })

  describe('static', () => {
    describe('#getUserSession', () => {
      const sessionUser = {
        username: 'Bob'
      }
      describe('SSR::when request is present', () => {
        let req
        beforeEach(() => {
          req = {
            session: {
              passport: {
                user: sessionUser
              }
            }
          }
        })

        it('returns false if no user session exists', async () => {
          delete req.session
          const userSession = await App.getUserSession(req)
          expect(userSession).toBe(false)
        })

        it('refreshes the user object and returns it', async () => {
          const userSession = await App.getUserSession(req)
          expect(OrgService).toHaveBeenCalledTimes(1)
          const mockOrgServiceInstance = OrgService.mock.instances[0]
          const mockRefreshUser = mockOrgServiceInstance.refreshUser
          expect(mockRefreshUser).toHaveBeenCalledWith(req.session.passport.user)
          expect(userSession).toEqual(req.session.passport.user)
        })

      })

      describe('CSR::no request present', () => {
        it('makes request to get user session', async () => {
          axios.get.mockResolvedValue({ data: sessionUser })

          const userSession = await App.getUserSession()
          expect(axios.get).toHaveBeenCalledTimes(1)
          expect(axios.get).toHaveBeenCalledWith(`${window.location.origin}/session/user`)
          expect(userSession).toEqual(userSession)
        })

        it('returns false if an error occurred', async () => {
          axios.get.mockRejectedValue('Some error')

          const userSession = await App.getUserSession()
          expect(userSession).toBe(false)
        })
      })
    })

    describe('#getUserTeamsDetails', () => {
      const allTeams = [{
        metadata: { name: 'hub-admin' },
        spec: { summary: 'Hub Admins' }
      }, {
        metadata: { name: 'team-a' },
        spec: { summary: 'Team A' }
      }, {
        metadata: { name: 'team-b' },
        spec: { summary: 'Team B' }
      }]

      it('returns detailed data for the users teams', async () => {
        apiRequest.mockResolvedValue({ items: allTeams })
        const userTeams = await App.getUserTeamsDetails(undefined, ['team-a', 'team-b'])
        expect(userTeams).toHaveLength(2)
        expect(userTeams[0].metadata.name).toBe('team-a')
        expect(userTeams[1].metadata.name).toBe('team-b')
      })

      it('filters out the hub-admin team', async () => {
        apiRequest.mockResolvedValue({ items: allTeams })
        const userTeams = await App.getUserTeamsDetails(undefined, ['hub-admin'])
        expect(userTeams).toHaveLength(0)
      })
    })

    describe('#getInitialProps', () => {
      const staticProps = {
        unrestrictedPage: true,
        otherProp: 'value'
      }
      const user = { username: 'bob' }
      const userTeams = { items: [] }
      const getUserSessionOriginal = App.getUserSession
      const getUserTeamsDetailsOriginal = App.getUserTeamsDetails

      beforeEach(() => {
        App.getUserSession = jest.fn().mockResolvedValue(user)
        App.getUserTeamsDetails = jest.fn().mockResolvedValue(userTeams)
      })

      afterEach(() => {
        App.getUserSession = getUserSessionOriginal
        App.getUserTeamsDetails = getUserTeamsDetailsOriginal
      })

      it('return props early if unrestrictedPage set as a staticProp as object', async () => {
        const Component = { staticProps }
        const props = await App.getInitialProps({ Component })
        expect(props).toEqual({ pageProps: staticProps })
      })

      it('can also read static props from function', async () => {
        const Component = {
          staticProps: () => staticProps
        }
        const props = await App.getInitialProps({ Component })
        expect(props).toEqual({ pageProps: staticProps })
      })

      it('redirects to login if session user is not found', async () => {
        App.getUserSession.mockResolvedValue(false)
        const params = {
          Component: {},
          ctx: {}
        }
        await App.getInitialProps(params)
        expect(redirect).toHaveBeenCalledTimes(1)
        expect(redirect).toHaveBeenCalledWith(undefined, '/login', true)
      })

      it('returns props, including pageProps, user and userTeams', async () => {
        const params = {
          Component: { staticProps: { myProp: 'myValue' }},
          ctx: {}
        }
        const props = await App.getInitialProps(params)
        expect(props).toEqual({
          pageProps: { myProp: 'myValue' },
          user,
          userTeams
        })
      })

      it('calls Component.getIntialProps and merges into pageProps', async () => {
        const initialProps = { prop1: 'hello', prop2: 'world' }
        const params = {
          Component: {
            staticProps: { myProp: 'myValue' },
            getInitialProps: jest.fn().mockResolvedValue(initialProps)
          },
          ctx: {}
        }
        const props = await App.getInitialProps(params)
        expect(params.Component.getInitialProps).toHaveBeenCalledTimes(1)
        expect(params.Component.getInitialProps).toHaveBeenCalledWith(params.ctx)
        expect(props).toEqual({
          pageProps: { myProp: 'myValue', ...initialProps },
          user,
          userTeams
        })
      })
    })
  })
})
