# Please set your repository in file:
- .golangci.yml

# Add secret in your Gihub
- GITHUB_TOKEN in dependabot_changelog
- CHANGELOG_PAT in generate_changelog.yml (to Push on main)
- DOC_PAT in page.yml (to generate documentation)

# Show example below:

<div align="center">
    <a href="https://github.com/FrangipaneTeam/terraform-plugin-framework-supertype/releases/latest">
      <img alt="Latest release" src="https://img.shields.io/github/v/release/FrangipaneTeam/terraform-plugin-framework-supertype?style=for-the-badge&logo=starship&color=C9CBFF&logoColor=D9E0EE&labelColor=302D41&include_prerelease&sort=semver" />
    </a>
    <a href="https://github.com/FrangipaneTeam/terraform-plugin-framework-supertype/pulse">
      <img alt="Last commit" src="https://img.shields.io/github/last-commit/FrangipaneTeam/terraform-plugin-framework-supertype?style=for-the-badge&logo=starship&color=8bd5ca&logoColor=D9E0EE&labelColor=302D41"/>
    </a>
    <a href="https://github.com/FrangipaneTeam/terraform-plugin-framework-supertype/stargazers">
      <img alt="Stars" src="https://img.shields.io/github/stars/FrangipaneTeam/terraform-plugin-framework-supertype?style=for-the-badge&logo=starship&color=c69ff5&logoColor=D9E0EE&labelColor=302D41" />
    </a>
    <a href="https://github.com/FrangipaneTeam/terraform-plugin-framework-supertype/issues">
      <img alt="Issues" src="https://img.shields.io/github/issues/FrangipaneTeam/terraform-plugin-framework-supertype?style=for-the-badge&logo=bilibili&color=F5E0DC&logoColor=D9E0EE&labelColor=302D41" />
    </a>
</div>

## terraform-plugin-framework-supertype

supertype is a custom type of Terraform type issue from Terraform schema, like SuperShema it's a meta Objet and function to permit you to manipulate Go object and Terrform object.

This is a try to solve these issues :

* CustomType is defined in the Schema on each Attributes fields on resources and datasources.
* Each kind of Terraform Type is take into account:
* * Simple: String, Bool, Int64, Float64
* * Array: List, Map, Set
* * Nested: ListNested, MapNested, SetNested, SingleNested


### Documentation

For more information about the supertype, please refer to the [documentation](https://github.frangipane.io/terraform/supertype/why).