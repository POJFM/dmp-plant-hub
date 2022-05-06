import { render, screen, fireEvent } from '@testing-library/react'
import '@testing-library/jest-dom'
import InitForm from './../components/InitForm'
import { ApolloProvider } from '@apollo/client'
import client from './../apollo/client'

test('Init logic', async () => {
	render(
		<ApolloProvider client={client}>
			<InitForm test={true} />
		</ApolloProvider>
	)

	global.XMLHttpRequest = undefined

	// limits on, scheduled off => everything apart from hour range is enabled

	expect(await screen.findByRole('button', { name: /uložit/i })).toBeEnabled()
	expect(screen.getByTestId('limitsTrigger')).toBeEnabled()
	expect(screen.getByTestId('scheduledTrigger')).toBeDisabled()
	expect(screen.getByTestId('irrigationDuration')).toBeEnabled()
	expect(screen.getByTestId('moistLimit')).toBeEnabled()
	expect(screen.getByTestId('waterAmountLimit')).toBeEnabled()
	expect(screen.getByTestId('waterLevelLimit')).toBeEnabled()
	expect(screen.getByTestId('hourRange')).toBeDisabled()
	expect(screen.getByTestId('location')).toBeEnabled()

	// limits off, scheduled on => moist limit, water amount limit and water level limit are disabled

	fireEvent.click(screen.getByTestId('limitsTrigger'))
	fireEvent.click(screen.getByTestId('scheduledTrigger'))

	expect(await screen.findByRole('button', { name: /uložit/i })).toBeEnabled()
	expect(screen.getByTestId('limitsTrigger')).toBeDisabled()
	expect(screen.getByTestId('scheduledTrigger')).toBeEnabled()
	expect(screen.getByTestId('irrigationDuration')).toBeEnabled()
	expect(screen.getByTestId('moistLimit')).toBeDisabled()
	expect(screen.getByTestId('waterAmountLimit')).toBeDisabled()
	expect(screen.getByTestId('waterLevelLimit')).toBeDisabled()
	expect(screen.getByTestId('hourRange')).toBeEnabled()
	expect(screen.getByTestId('location')).toBeEnabled()

	// limits off, scheduled off => irrigationDuration and location are enabled

	fireEvent.click(screen.getByTestId('scheduledTrigger'))

	expect(await screen.findByRole('button', { name: /uložit/i })).toBeDisabled()
	expect(screen.getByTestId('limitsTrigger')).toBeDisabled()
	expect(screen.getByTestId('scheduledTrigger')).toBeDisabled()
	expect(screen.getByTestId('irrigationDuration')).toBeEnabled()
	expect(screen.getByTestId('moistLimit')).toBeDisabled()
	expect(screen.getByTestId('waterAmountLimit')).toBeDisabled()
	expect(screen.getByTestId('waterLevelLimit')).toBeDisabled()
	expect(screen.getByTestId('hourRange')).toBeDisabled()
	expect(screen.getByTestId('location')).toBeEnabled()

	// limits on, scheduled on => everything is enabled

	fireEvent.click(screen.getByTestId('limitsTrigger'))
	fireEvent.click(screen.getByTestId('scheduledTrigger'))

	expect(await screen.findByRole('button', { name: /uložit/i })).toBeEnabled()
	expect(screen.getByTestId('limitsTrigger')).toBeEnabled()
	expect(screen.getByTestId('scheduledTrigger')).toBeEnabled()
	expect(screen.getByTestId('irrigationDuration')).toBeEnabled()
	expect(screen.getByTestId('moistLimit')).toBeEnabled()
	expect(screen.getByTestId('waterAmountLimit')).toBeEnabled()
	expect(screen.getByTestId('waterLevelLimit')).toBeEnabled()
	expect(screen.getByTestId('hourRange')).toBeEnabled()
	expect(screen.getByTestId('location')).toBeEnabled()
})
