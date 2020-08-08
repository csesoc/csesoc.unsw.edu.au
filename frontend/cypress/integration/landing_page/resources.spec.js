describe('Resources Landing Page', () => {
  beforeEach(() => {
    cy.visit('/')
  })
  
  it('finds whether the resources are displayed', () => {
    cy
      .get('[data-cy=preview-title]')
      .should('have.length.gt', 1)
    cy
      .get('[data-cy=preview-description]')
      .should('have.length.gt', 1)
  });

  it('finds an image pre-displayed', () => {
    cy
      .get('[data-cy=preview-image]')
      .should('be.visible')
  })
})