describe('Community Links Landing Page', () => {
  beforeEach(() => {
    cy.visit('/');
  });

  it('finds whether the text and images are displayed', () => {
    // Making sure that the community links are included in the landing page
    cy.get('[data-cy=communityLinks]');
    // making sure that the body text is there
    cy.get('[data-cy=communityLinks-images]');

    // making sure that each image is being displayed, and is visible after hovering
    // TODO: figure out how to test the zoom feature, and also add sizing into it aswell
    cy.get('[data-cy=communityLinks-images]')
      .each(($logo) => {
        $logo.trigger('mouseover');
      });
  });
});
