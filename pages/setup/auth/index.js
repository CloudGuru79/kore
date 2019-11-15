import Link from 'next/link'
import { Button, Row, Col, Card, Typography } from 'antd'
const { Title } = Typography;

const SetupAuthIndexPage = () => (
  <div style={{ background: '#fff', padding: 24, minHeight: 280 }}>
    <Row type="flex" justify="center">
      <Col span={24}>
        <Card style={{ backgroundColor: '#f5f5f5' }}>
          <Title>Welcome to the Hub!</Title>
          <Title level={4}>Authentication is not currently configured</Title>
          <p>This must be done before any teams can start using the hub</p>
          <Button type="primary">
            <Link href="/setup/auth/configure">
              <a>Begin setup</a>
            </Link>
          </Button>
        </Card>
      </Col>
    </Row>
  </div>
)

SetupAuthIndexPage.getInitialProps = async () => {
  return {
    title: 'No authentication provider configured',
    unrestrictedPage: true
  }
}

export default SetupAuthIndexPage
