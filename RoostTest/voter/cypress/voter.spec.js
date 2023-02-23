

describe('Visit ROOST_SVC_URL and check the title', () => {
  it('should visit ROOST_SVC_URL and check the title', () => {
    cy.visit(Cypress.env('ROOST_SVC_URL'))
    cy.title().should('eq', 'Roost')
  })
})