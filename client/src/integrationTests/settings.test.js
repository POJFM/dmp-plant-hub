import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import Settings from './../components/Settings'
import { ApolloProvider } from '@apollo/client'
import client from './../apollo/client'

test('Settings logic', async () => {
	render(
		<ApolloProvider client={client}>
			<Settings />
		</ApolloProvider>
	)
})
