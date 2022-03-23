import { ApolloClient, InMemoryCache } from '@apollo/client'

const client = new ApolloClient({
	uri: `${process.env.REACT_APP_GO_API_URL}/query`,
	cache: new InMemoryCache(),
})

export default client
