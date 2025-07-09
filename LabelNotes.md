**Subject: Open Source System Alternatives for Processing ZPL Label Data on Local Servers (On-Premise)**

In our company's operational processes for creating ZPL-based labels, the currently used external web services (such as Labelary.com) pose risks in terms of data privacy and information security. To prevent sensitive information such as customer or product data from being sent to servers outside our control, it is aimed to transition to a solution that will host this functionality in our own infrastructure (on-premise).

As a result of research conducted for this need, the following reliable and open-source alternatives that can convert ZPL data to formats such as PNG/PDF have been identified:

1.  **zpl-tool:** A complete solution that offers a modern web interface similar to Labelary, can be easily installed with Docker, and includes advanced features such as live preview. It is the most capable alternative that can be quickly deployed.

    - **Repo Link:** [https://github.com/enoy19/zpl-tool](https://github.com/enoy19/zpl-tool)

    DeveloperNote:

    - Tested locally. Works with Docker.
    - When working with Docker, I keep the binarykits-zpl service enabled in the docker-compose.yml file.
    - Has an interface written with SvelteKit. Can be manipulated (I tried it).
    - Forking the repo is sufficient. Not many stars but made me think.

2.  **Zebrash:** A high-performance Go library developed as a direct alternative to Labelary. It is quite suitable for adding ZPL to image generation capability to our existing systems as a microservice.

    - **Repo Link:** [https://github.com/ingridhq/zebrash](https://github.com/ingridhq/zebrash)

    DeveloperNote:

    - Works with Docker.
    - Forking the repo is sufficient. Has a bit more stars.
    - No ready UI, I added it.

3.  **ZplDesigner:** A C#/.NET-based ZPL processing library. It offers a strategic advantage as it can be easily integrated into existing .NET-based projects in our company.
    - **Repo Link:** [https://github.com/IkeRolfe/ZplDesigner](https://github.com/IkeRolfe/ZplDesigner)

Adopting one of these projects will increase our data security while eliminating external dependencies and provide cost advantages in the long term.
