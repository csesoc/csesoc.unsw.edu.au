describe('Resources Landing Page', () => {
  beforeEach(() => {
    cy.visit('/');
  });

  it('finds whether the text and images are displayed', () => {
    // Make sure that all 4 boxes appear
    cy.get('[data-cy=resources-box]').should('have.length', 4);

    // Make sure that all 4 boxes appear
    cy.get('[data-cy=resources-img]').each(($image) => {
      // Check if transform transition is defined
      cy.wrap($image).should('have.css', 'transition').and('eq', 'transform 0.2s ease 0s');
    });
  });
});
