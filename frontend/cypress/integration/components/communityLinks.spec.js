describe('Community Links Landing Page', () => {
  beforeEach(() => {
    cy.visit('/');
  });

  it('finds whether the text and images are displayed', () => {
    // Making sure that the community links are included in the landing page
    cy.get('[data-cy=communityLinks]');
    // making sure that the join us text is there
    cy.get('[data-cy=communityLinks-title]');
    // making sure that the body text is there
    cy.get('[data-cy=communityLinks-bodyText]');
    // making sure that the body text is there
    cy.get('[data-cy=communityLinks-images]');

    // making sure that each image is being displayed, and is visible after hovering
    // TODO: figure out how to test the zoom feature, and also add sizing into it aswell
    cy.get('[data-cy=communityLinks-facebookImage].zoom')
      .trigger('mouseover');
    cy.get('[data-cy=communityLinks-slackImage].zoom')
      .trigger('mouseover');
    cy.get('[data-cy=communityLinks-discordImage].zoom')
      .trigger('mouseover');
  });
});
