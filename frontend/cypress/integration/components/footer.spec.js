describe("Footer Testing", () => {
    beforeEach(() => {
        cy.visit('/');
    });

    it('finds whether the logo links to the homepage and the address is displayed', () => {
        // Check if the logo has link
        cy.get('[data-cy=footer-main-logo]').should('have.attr', 'href', 'localhost:8080/');

        // Check if the address is displayed
        cy.get('[data-cy=footer-address]').contains('B03 CSE Building K17, UNSW');

        // Check if the internal link is displayed

        cy.get('[data-cy=footer-internal-link]').should( ($lis) => {
            expect($lis, '5 items').to.have.length(5);
            // Check if the link is correctly linked
            expect($lis.eq(0), 'first item').to.have.attr('href', 'localhost:8080/');
            // Check if the text is correctly displayed
            expect($lis.eq(0), 'first item').to.have.text('Home');
            // Check if the link is correctly linked
            expect($lis.eq(1), 'second item').to.have.attr('href', 'localhost:8080/about');
            // Check if the text is correctly displayed
            expect($lis.eq(1), 'second item').to.have.text('About');
            // Check if the link is correctly linked
            expect($lis.eq(2), 'third item').to.have.attr('href', 'localhost:8080/resources');
            // Check if the text is correctly displayed
            expect($lis.eq(2), 'third item').to.have.text('Resources');
            // Check if the link is correctly linked
            expect($lis.eq(3), 'forth item').to.have.attr('href', 'localhost:8080/sponsor');
            // Check if the text is correctly displayed
            expect($lis.eq(3), 'forth item').to.have.text('Sponsor');
            // Check if the link is correctly linked
            expect($lis.eq(4), 'fifth item').to.have.attr('href', 'localhost:8080/engage');
            // Check if the text is correctly displayed
            expect($lis.eq(4), 'fifth item').to.have.text('Engage');
        });

        // Check if the external link is displayed and the logo is displayed

        cy.get('[data-cy=footer-social-media-link]').should(($lis) => {
            expect($lis, '8 items').to.have.length(8);
            expect($lis.eq(0), 'first item').to.have.attr('href', 'https://www.facebook.com/csesoc');
            expect($lis.eq(0), 'first item').to.have.text('Facebook Page');
            expect($lis.eq(1), 'second item').to.have.attr('href', 'https://www.facebook.com/csesoc');
            expect($lis.eq(1), 'second item').to.have.text('Facebook Page');

        });
    });
})