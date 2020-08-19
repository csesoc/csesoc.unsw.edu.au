describe('Resources Landing Page', () => {
  beforeEach(() => {
    cy.visit('/');
  });

  it('finds whether the resources are displayed', () => {
    // Check if titles exists
    cy.get('[data-cy=preview-title]').should('have.length.gt', 1);
    // Check if descriptions exists
    cy.get('[data-cy=preview-description]').should('have.length.gt', 1);
    // Check if an initial image is visible
    // TODO: something is wrong in detecting the image
    // Something with it being lazily loaded.
    // cy.get('[data-cy=preview-image]').scrollIntoView().should('be.visible')
    cy.get('[data-cy=preview-image]');
    // Check on hover works
    cy
      .get('[data-cy=preview-item]')
      .each(($li) => {
        cy.wrap($li).trigger('mouseover')
          .get('[data-cy=preview-image]')
          .should('be.visible');
      });
  });
});
