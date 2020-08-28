describe('Community Links Landing Page', () => {
    beforeEach(() => {
      cy.visit('/');
    });
  
    it('finds whether the text and images are displayed', () => {
      // TODO: fix padding according to style scape and add to testing

      // Making sure that the community links are included in the landing page
      cy.get('[data-cy=communitLinks]');
      // making sure that the join us text is there
      cy.get('[data-cy=communityLinks-title]');
      // making sure that the body text is there
      cy.get('[data-cy=communityLinks-bodyText]');
      // making sure that the body text is there
      cy.get('[data-cy=communityLinks-joinCommunityText]');
      // making sure that the body text is there
      cy.get('[data-cy=communityLinks-images]');

      // making sure that each image is being displayed, and is visible after hovering
      // TODO: figure out how to test the zoom feature, and also add sizing into it aswell
      cy.get('[data-cy=communityLinks-facebookImage].zoom')
        .trigger('mouseover')
      cy.get('[data-cy=communityLinks-slackImage].zoom')
        .trigger('mouseover')
      cy.get('[data-cy=communityLinks-discordImage].zoom')
        .trigger('mouseover')
    });

    it('tests scrolling to the join us section from the top of the landing page', () => {
      // finding the join us button and clicking it

      // cypress is giving a 'the chainer inViewPort cannot be found', 
      // might try making it a command not an assert to see if that fixes it
      // ref: https://github.com/cypress-io/cypress/issues/877
      cy.get('[data-cy=communitLinks]').should('not.be.inViewPort');
      cy.get('[data-cy=joinus-button]').click();
      // it should scroll to the join us section
      cy.get('[data-cy=communitLinks]').should('not.be.inViewPort');
      
    });

  });
  