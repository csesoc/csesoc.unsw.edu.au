describe("Footer Testing", () => {
    beforeEach(() => {
        cy.visit('/');
    });

    it('finds whether the logo links to the homepage and the address is displayed', () => {
        // Check if the logo has link
        cy.get('[data-cy=footer-main-logo-link]').should('have.attr', 'href', '/');
        cy.get('[data-cy=footer-main-logo]');

        // Check if the address is displayed
        cy.get('[data-cy=footer-address]').contains('B03 CSE Building K17, UNSW');

        // Check if the internal link is displayed
        cy.get('[data-cy=footer-internal-link]').should( ($lis) => {
            // Check there is the correct number of items
            expect($lis, '5 items').to.have.length(5);
            // Check if the link is correctly linked
            expect($lis.eq(0), 'first item').to.have.attr('href', '/');
            // Check if the text is correctly displayed
            expect($lis.eq(0), 'first item').to.have.text('Home');
            // Check if the link is correctly linked
            expect($lis.eq(1), 'second item').to.have.attr('href', '/#/about');
            // Check if the text is correctly displayed
            expect($lis.eq(1), 'second item').to.have.text('About');
            // Check if the link is correctly linked
            expect($lis.eq(2), 'third item').to.have.attr('href', '/#/resources');
            // Check if the text is correctly displayed
            expect($lis.eq(2), 'third item').to.have.text('Resources');
            // Check if the link is correctly linked
            expect($lis.eq(3), 'forth item').to.have.attr('href', '/#/sponsor');
            // Check if the text is correctly displayed
            expect($lis.eq(3), 'forth item').to.have.text('Sponsor');
            // Check if the link is correctly linked
            expect($lis.eq(4), 'fifth item').to.have.attr('href', '/#/engage');
            // Check if the text is correctly displayed
            expect($lis.eq(4), 'fifth item').to.have.text('Engage');
        });

        // Check if the address is displayed
        cy.get('[data-cy=footer-media-title]').contains('Social Media');

        // Check if the internal link is displayed
        cy.get('[data-cy=footer-social-media-link]').should(($lis) => {
            // Check there is the correct number of items
            expect($lis, '8 items').to.have.length(8);
            // Check if the link is correctly linked
            expect($lis.eq(0), 'first item').to.have.attr('href', 'https://www.facebook.com/csesoc');
            // Check if the text is correctly displayed
            expect($lis.eq(0), 'first item').to.have.text('Facebook Page');
            // Check if the link is correctly linked
            expect($lis.eq(1), 'second item').to.have.attr('href', 'https://www.facebook.com/groups/csesoc');
            // Check if the text is correctly displayed
            expect($lis.eq(1), 'second item').to.have.text('Facebook Group');
            // Check if the link is correctly linked
            expect($lis.eq(2), 'third item').to.have.attr('href', 'https://www.instagram.com/csesoc_unsw');
            // Check if the text is correctly displayed
            expect($lis.eq(2), 'third item').to.have.text('Instagram');
            // Check if the link is correctly linked
            expect($lis.eq(3), 'forth item').to.have.attr('href', 'https://forms.office.com/Pages/ResponsePage.aspx?id=pM_2PxXn20i44Qhnufn7o6ecLZTBorREjnXuTY-JfmBUMEpOMFBDTU1UWkhBWllWRTNPOVJFMUNCRi4u');
            // Check if the text is correctly displayed
            expect($lis.eq(3), 'forth item').to.have.text('Discord Community');
            // Check if the link is correctly linked
            expect($lis.eq(4), 'fifth item').to.have.attr('href', 'https://csesoc-community.slack.com/');
            // Check if the text is correctly displayed
            expect($lis.eq(4), 'fifth item').to.have.text('Slack Community');
            // Check if the link is correctly linked
            expect($lis.eq(5), 'sixth item').to.have.attr('href', 'https://www.linkedin.com/company/csesoc/');
            // Check if the text is correctly displayed
            expect($lis.eq(5), 'sixth item').to.have.text('LinkedIn');
            // Check if the link is correctly linked
            expect($lis.eq(6), 'seventh item').to.have.attr('href', '#');
            // Check if the text is correctly displayed
            expect($lis.eq(6), 'seventh item').to.have.text('Tiktok');
            // Check if the link is correctly linked
            expect($lis.eq(7), 'eighth item').to.have.attr('href', 'https://www.youtube.com/channel/UC1JHpRrf9j5IKluzXhprUJg');
            // Check if the text is correctly displayed
            expect($lis.eq(7), 'eighth item').to.have.text('YouTube');
        });

        // Check if the address is displayed
        cy.get('[data-cy=footer-resources-title]').contains('For Your Better Future');

        // Check if the resources link is displayed
        cy.get('[data-cy=footer-resources-link]').should(($lis) => {
            // Check there is the correct number of items
            expect($lis, '5 items').to.have.length(5);
            // Check if the link is correctly linked
            expect($lis.eq(0), 'first item').to.have.attr('href', 'https://media.csesoc.org.au/');
            // Check if the text is correctly displayed
            expect($lis.eq(0), 'first item').to.have.text('CSESoc Media');
            // Check if the link is correctly linked
            expect($lis.eq(1), 'second item').to.have.attr('href', 'https://blog.csesoc.org.au/');
            // Check if the text is correctly displayed
            expect($lis.eq(1), 'second item').to.have.text('CSESoc Blog');
            // Check if the link is correctly linked
            expect($lis.eq(2), 'third item').to.have.attr('href', 'https://compclub.csesoc.unsw.edu.au/');
            // Check if the text is correctly displayed
            expect($lis.eq(2), 'third item').to.have.text('CSESoc Compclub');
            // Check if the link is correctly linked
            expect($lis.eq(3), 'forth item').to.have.attr('href', 'https://www.engineering.unsw.edu.au/computer-science-engineering/');
            // Check if the text is correctly displayed
            expect($lis.eq(3), 'forth item').to.have.text('UNSW CSE');
            // Check if the link is correctly linked
            expect($lis.eq(4), 'fifth item').to.have.attr('href', 'https://www.handbook.unsw.edu.au/');
            // Check if the text is correctly displayed
            expect($lis.eq(4), 'fifth item').to.have.text('UNSW Handbook');
        });
    });
})