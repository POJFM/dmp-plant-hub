import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import InitForm from './../components/InitForm'
import { ApolloProvider } from '@apollo/client'
import client from './../apollo/client'

test('Init logic', async () => {
	render(
		<ApolloProvider client={client}>
			<InitForm />
		</ApolloProvider>
	)

	// limits on, scheduled off => everything apart from hour range is enabled
	// limits off, scheduled on => moist limit, water amount limit and water level limit are disabled
	// limits off, scheduled off => everything is disabled
	// limits on, scheduled on => everything is enabled

	// expect(await screen.findByRole('button', { name: /initsave/i })).toBeDisabled()
	// expect(screen.findByRole('checkbox', { name: /limitsTrigger/i })).toBeEnabled()
	// expect(screen.findByRole('checkbox', { name: /scheduledTrigger/i })).toBeDisabled()

	// userEvent.type(screen.getByPlaceholderText(/amount/i), "50");
	// userEvent.type(screen.getByPlaceholderText(/add a note/i), "dinner");

	// expect(await screen.findByRole("button", { name: /pay/i })).toBeEnabled();
})
