import React from 'react'
import axios from 'axios'
import PropTypes from 'prop-types'

import apiRequest from '../../../lib/utils/api-request'
import apiPaths from '../../../lib/utils/api-paths'
import Breadcrumb from '../../../lib/components/Breadcrumb'
import AuditViewer from '../../../lib/components/AuditViewer'

class TeamAuditPage extends React.Component {
  static propTypes = {
    team: PropTypes.object.isRequired,
    events: PropTypes.array.isRequired,
  }

  state = {
    events: []
  }

  static staticProps = {
    title: 'Team Audit Viewer',
    adminOnly: false
  }

  static async getPageData({ req, res, query }) {
    const name = query.name
    const getTeam = () => apiRequest({ req, res }, 'get', apiPaths.team(name).self)
    const getAuditEvents = () => apiRequest({ req, res }, 'get', apiPaths.team(name).audit)

    return axios.all([getTeam(), getAuditEvents()])
      .then(axios.spread(function (team, eventList) {
        return { team, events: eventList.items }
      }))
      .catch(err => {
        throw new Error(err.message)
      })
  }

  static getInitialProps = async ctx => {
    const data = await TeamAuditPage.getPageData(ctx)
    return data
  }

  constructor(props) {
    super(props)
    this.state = { events: props.events }
  }

  render() {
    const teamName = this.props.team.metadata.name

    return (
      <div>
        <Breadcrumb
          items={[
            { text: this.props.team.spec.summary, href: '/teams/[name]', link: `/teams/${teamName}` },
            { text: 'Team Audit Viewer' }
          ]}
        />
        <AuditViewer items={this.state.events} />
      </div>
    )
  }
}

export default TeamAuditPage