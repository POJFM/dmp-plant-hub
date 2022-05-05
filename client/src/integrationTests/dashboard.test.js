import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import Dashboard from './../components/Dashboard'
import { ApolloProvider } from '@apollo/client'
import client from './../apollo/client'

test('Dashboard logic', async () => {
	render(
		<ApolloProvider client={client}>
			<Dashboard />
		</ApolloProvider>
	)
})
