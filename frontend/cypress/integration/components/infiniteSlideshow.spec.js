describe('InfiniteSlideshow Landing Page', () => {
  beforeEach(() => {
    cy.visit('/');
  });

  it('finds whether the banners are displayed', () => {
    // Check if banners exist on each slot
    cy.get('[data-cy=infinite-slideshow]').each(($is) => {
      // Check if it is animated
      cy.wrap($is).should('have.css', 'animation-name');
      // Check if the animation is infinite
      cy.wrap($is).should('have.css', 'animation-iteration-count').and('eq', 'infinite');
      // Check if the animation is running
      cy.wrap($is).should('have.css', 'animation-play-state').and('eq', 'running');
      // Each slot should have at least 1 item in it
      cy.wrap($is).children().each(($slot) => {
        cy.wrap($slot).children().should('have.length.gt', 0);
      });
    });
  });
});
