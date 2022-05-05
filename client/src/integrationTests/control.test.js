import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import Control from './../components/Control'
import { ApolloProvider } from '@apollo/client'
import client from './../apollo/client'

test('Control logic', async () => {
	render(
		<ApolloProvider client={client}>
			<Control />
		</ApolloProvider>
	)
})
