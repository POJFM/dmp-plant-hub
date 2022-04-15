import { ApolloClient, InMemoryCache } from '@apollo/client'

const client = new ApolloClient({
	uri: `http://${process.env.REACT_APP_GO_IP}:5000/query`,
	cache: new InMemoryCache(),
})

export default client
