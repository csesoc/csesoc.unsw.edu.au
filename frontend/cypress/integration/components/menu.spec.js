describe('Menu', () => {
  beforeEach(() => {
    cy.visit('/');
  });

  it('navigates to all internal website pages', () => {
    // Open the menu
    cy.get('[data-cy=menu-toggle]').click();
    // Navigate to About page
    cy.get('[data-cy=about-link]').click();
    // Check current url
    cy.url().should('match', /about$/);

    // Open the menu
    cy.get('[data-cy=menu-toggle]').click();
    // Navigate to Resources page
    cy.get('[data-cy=resources-link]').click();
    // Check current url
    cy.url().should('match', /resources$/);

    // Open the menu
    cy.get('[data-cy=menu-toggle]').click();
    // Navigate to Sponsors page
    cy.get('[data-cy=sponsors-link]').click();
    // Check current url
    cy.url().should('match', /sponsors$/);

    // Open the menu
    cy.get('[data-cy=menu-toggle]').click();
    // Navigate to Engage page
    cy.get('[data-cy=engage-link]').click();
    // Check current url
    cy.url().should('match', /engage$/);
  });

  it('checks all external links', () => {
    // Open the menu
    cy.get('[data-cy=menu-toggle]').click();

    // Check all the social links are present and have a href field
    cy.get('[data-cy=menu-social-link] > a').should('have.length', 7).each(($link) => {
      cy.wrap($link).should('have.attr', 'href').should('not.be.empty');
    });
  });
});
