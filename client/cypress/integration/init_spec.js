desribe('init', () => {
	cy.visit('/')
  cy.findByRole('checkbox', {name: /limitsTrigger/i})

  // limits on, scheduled off => everything apart from hour range is enabled
  // limits off, scheduled on => moist limit, water amount limit and water level limit are disabled
  // limits off, scheduled off => everything is disabled
  // limits on, scheduled on => everything is enabled
})
