# Changelog

## [2.0.0](https://github.com/devsy-org/api/compare/v1.0.0...v2.0.0) (2026-04-18)


### ⚠ BREAKING CHANGES

* API group renamed from loft.sh to devsy.sh. All Loft and VCluster type names renamed to Devsy equivalents.
* rename module from loft-sh/api/v4 to skevetter/api ([#11](https://github.com/devsy-org/api/issues/11))
* rename module from loft-sh/api/v4 to skevetter/api ([#10](https://github.com/devsy-org/api/issues/10))

### Features

* add annotation to exclude clusters for cluster account templates ([4a13959](https://github.com/devsy-org/api/commit/4a139590f6a067930422ecf0becd81c05dd510bc))
* add cluster account template frontend ([087dc69](https://github.com/devsy-org/api/commit/087dc69829e1c089ddad2f94f7880fd4812d6a15))
* add cluster reset for rbac & kiosk ([4e693a4](https://github.com/devsy-org/api/commit/4e693a4d5ff9c5cf6b97d03eb7f63277b8c8ccab))
* add profile picture ([1398c6a](https://github.com/devsy-org/api/commit/1398c6a03d20b96a2cb3f24191b6edc7b69f31b4))
* add sleepmode.loft.sh/delete-after ([0a96da7](https://github.com/devsy-org/api/commit/0a96da7dfe6e4d734648c095fbbbd497b291e6ae))
* add spaces sub view ([6e9c2ec](https://github.com/devsy-org/api/commit/6e9c2ec8505d4c4d0dd05d6c55a8940f2528dba7))
* add subdomain backend ([7651a1a](https://github.com/devsy-org/api/commit/7651a1a3fbd74537948e2d224db6613a7d13c6a0))
* chart username & password default ([7a4d167](https://github.com/devsy-org/api/commit/7a4d16752c8c096dcd00c13822843347f782bc88))
* cluster select frontend & antdesign update ([2844fbc](https://github.com/devsy-org/api/commit/2844fbcb94d35b8c64ab8956546adccad3aca9be))
* disable / enable users & token generation ([8bab235](https://github.com/devsy-org/api/commit/8bab235ad61f3e39ea2b2de4912ac83a1c2cc7b8))
* force sleep duration ([c4b5137](https://github.com/devsy-org/api/commit/c4b5137ac72123e17df2fc9236280c184b950962))
* image pull secrets ui ([e430a1c](https://github.com/devsy-org/api/commit/e430a1cb7b07a5098241a16cb956ff9c666342c6))
* improve loft login ([df87e59](https://github.com/devsy-org/api/commit/df87e59443f94c16a21035fd24fc36ff60c8248a))
* make air-gapped installation easier ([950171b](https://github.com/devsy-org/api/commit/950171bcdaf9a885eaa7b50cb9d404b1712ef0e3))
* **oidc:** add auth.oidc.insecureCa & auth.oidc.groups ([0863db8](https://github.com/devsy-org/api/commit/0863db83f87b64d29b03d333abdf046d63d602e3))
* **oidc:** implement support for thin oidc tokens ([8849a51](https://github.com/devsy-org/api/commit/8849a51412f5a51533e4688cc6458701f7a93bf4))
* rename loft/vcluster to devsy ([#14](https://github.com/devsy-org/api/issues/14)) ([c42a044](https://github.com/devsy-org/api/commit/c42a044a69cb358d6bad6cc89b51c8ea0a0b7889))
* rename module from loft-sh/api/v4 to skevetter/api ([#10](https://github.com/devsy-org/api/issues/10)) ([3241f19](https://github.com/devsy-org/api/commit/3241f197a59e36c90bb2490cade2f21b73a09669))
* rename module from loft-sh/api/v4 to skevetter/api ([#11](https://github.com/devsy-org/api/issues/11)) ([e2b7388](https://github.com/devsy-org/api/commit/e2b73888d90b4137fd8383c328bafa99d90313c7))
* restart loft automatically on config change & config comments ([bc7866e](https://github.com/devsy-org/api/commit/bc7866e7b64c23f81af95fb0a9eae5a67722b641))
* secrets frontend & crd ([c5e8b94](https://github.com/devsy-org/api/commit/c5e8b94adc6a1a64aff40c0ac66b1bcd133a524b))
* skip insecure tls verify helm & helm update ([1da9f07](https://github.com/devsy-org/api/commit/1da9f07129b2b6f945849f6e3567e0b470209100))
* sync cluster account templates ([cc416f2](https://github.com/devsy-org/api/commit/cc416f27b673c6800e04a7470b46ddcb059f0467))
* update deps to devsy-org and remove legacy references ([#1](https://github.com/devsy-org/api/issues/1)) ([40182d4](https://github.com/devsy-org/api/commit/40182d42f6cfdbb9e3b05f65a1b8b23b82fd03ae))


### Bug Fixes

* add MegaLinter config and upgrade to v9 ([#17](https://github.com/devsy-org/api/issues/17)) ([17ddaaf](https://github.com/devsy-org/api/commit/17ddaaf08f767bbcb1a2d60298c2a2d818fd0944))
* fixes an issue where the profile picture is changed all the time ([846a0a1](https://github.com/devsy-org/api/commit/846a0a19a31f2939602352ee82b4fe4e6de2c3b7))
* fixes an issue with user & team api migration ([1b38d0c](https://github.com/devsy-org/api/commit/1b38d0c675cbf42d3f070e0c4769ef9ec899061f))
* fixes e2e tests ([7ddeec8](https://github.com/devsy-org/api/commit/7ddeec8d54f5e37a5aa6fbcece0eb44d26894cd7))
* remove non existent secret client ([992039e](https://github.com/devsy-org/api/commit/992039e2b9bd0265926a54ae8fdce0422c0e4a96))
* set release-please manifest to v1.0.0 ([#15](https://github.com/devsy-org/api/issues/15)) ([c16f14b](https://github.com/devsy-org/api/commit/c16f14b307a8fd8433d3ecfd4772456c2f2cfac6))
* update versions to latest tag ([64145fb](https://github.com/devsy-org/api/commit/64145fb02515d386eaebfd7ccf596a5c8e80d56d))
* update versions to latest tag ([68ba5fb](https://github.com/devsy-org/api/commit/68ba5fb484c9d81f0e765fb65b797c965521b360))

## 1.0.0 (2026-04-13)


### ⚠ BREAKING CHANGES

* API group renamed from loft.sh to devsy.sh. All Loft and VCluster type names renamed to Devsy equivalents.
* rename module from loft-sh/api/v4 to skevetter/api ([#11](https://github.com/devsy-org/api/issues/11))
* rename module from loft-sh/api/v4 to skevetter/api ([#10](https://github.com/devsy-org/api/issues/10))

### Features

* add annotation to exclude clusters for cluster account templates ([4a13959](https://github.com/devsy-org/api/commit/4a139590f6a067930422ecf0becd81c05dd510bc))
* add cluster account template frontend ([087dc69](https://github.com/devsy-org/api/commit/087dc69829e1c089ddad2f94f7880fd4812d6a15))
* add cluster reset for rbac & kiosk ([4e693a4](https://github.com/devsy-org/api/commit/4e693a4d5ff9c5cf6b97d03eb7f63277b8c8ccab))
* add profile picture ([1398c6a](https://github.com/devsy-org/api/commit/1398c6a03d20b96a2cb3f24191b6edc7b69f31b4))
* add sleepmode.loft.sh/delete-after ([0a96da7](https://github.com/devsy-org/api/commit/0a96da7dfe6e4d734648c095fbbbd497b291e6ae))
* add spaces sub view ([6e9c2ec](https://github.com/devsy-org/api/commit/6e9c2ec8505d4c4d0dd05d6c55a8940f2528dba7))
* add subdomain backend ([7651a1a](https://github.com/devsy-org/api/commit/7651a1a3fbd74537948e2d224db6613a7d13c6a0))
* chart username & password default ([7a4d167](https://github.com/devsy-org/api/commit/7a4d16752c8c096dcd00c13822843347f782bc88))
* cluster select frontend & antdesign update ([2844fbc](https://github.com/devsy-org/api/commit/2844fbcb94d35b8c64ab8956546adccad3aca9be))
* disable / enable users & token generation ([8bab235](https://github.com/devsy-org/api/commit/8bab235ad61f3e39ea2b2de4912ac83a1c2cc7b8))
* force sleep duration ([c4b5137](https://github.com/devsy-org/api/commit/c4b5137ac72123e17df2fc9236280c184b950962))
* image pull secrets ui ([e430a1c](https://github.com/devsy-org/api/commit/e430a1cb7b07a5098241a16cb956ff9c666342c6))
* improve loft login ([df87e59](https://github.com/devsy-org/api/commit/df87e59443f94c16a21035fd24fc36ff60c8248a))
* make air-gapped installation easier ([950171b](https://github.com/devsy-org/api/commit/950171bcdaf9a885eaa7b50cb9d404b1712ef0e3))
* **oidc:** add auth.oidc.insecureCa & auth.oidc.groups ([0863db8](https://github.com/devsy-org/api/commit/0863db83f87b64d29b03d333abdf046d63d602e3))
* **oidc:** implement support for thin oidc tokens ([8849a51](https://github.com/devsy-org/api/commit/8849a51412f5a51533e4688cc6458701f7a93bf4))
* rename loft/vcluster to devsy ([#14](https://github.com/devsy-org/api/issues/14)) ([c42a044](https://github.com/devsy-org/api/commit/c42a044a69cb358d6bad6cc89b51c8ea0a0b7889))
* rename module from loft-sh/api/v4 to skevetter/api ([#10](https://github.com/devsy-org/api/issues/10)) ([3241f19](https://github.com/devsy-org/api/commit/3241f197a59e36c90bb2490cade2f21b73a09669))
* rename module from loft-sh/api/v4 to skevetter/api ([#11](https://github.com/devsy-org/api/issues/11)) ([e2b7388](https://github.com/devsy-org/api/commit/e2b73888d90b4137fd8383c328bafa99d90313c7))
* restart loft automatically on config change & config comments ([bc7866e](https://github.com/devsy-org/api/commit/bc7866e7b64c23f81af95fb0a9eae5a67722b641))
* secrets frontend & crd ([c5e8b94](https://github.com/devsy-org/api/commit/c5e8b94adc6a1a64aff40c0ac66b1bcd133a524b))
* skip insecure tls verify helm & helm update ([1da9f07](https://github.com/devsy-org/api/commit/1da9f07129b2b6f945849f6e3567e0b470209100))
* sync cluster account templates ([cc416f2](https://github.com/devsy-org/api/commit/cc416f27b673c6800e04a7470b46ddcb059f0467))


### Bug Fixes

* add MegaLinter config and upgrade to v9 ([#17](https://github.com/devsy-org/api/issues/17)) ([17ddaaf](https://github.com/devsy-org/api/commit/17ddaaf08f767bbcb1a2d60298c2a2d818fd0944))
* fixes an issue where the profile picture is changed all the time ([846a0a1](https://github.com/devsy-org/api/commit/846a0a19a31f2939602352ee82b4fe4e6de2c3b7))
* fixes an issue with user & team api migration ([1b38d0c](https://github.com/devsy-org/api/commit/1b38d0c675cbf42d3f070e0c4769ef9ec899061f))
* fixes e2e tests ([7ddeec8](https://github.com/devsy-org/api/commit/7ddeec8d54f5e37a5aa6fbcece0eb44d26894cd7))
* remove non existent secret client ([992039e](https://github.com/devsy-org/api/commit/992039e2b9bd0265926a54ae8fdce0422c0e4a96))
* set release-please manifest to v1.0.0 ([#15](https://github.com/devsy-org/api/issues/15)) ([c16f14b](https://github.com/devsy-org/api/commit/c16f14b307a8fd8433d3ecfd4772456c2f2cfac6))
* update versions to latest tag ([64145fb](https://github.com/devsy-org/api/commit/64145fb02515d386eaebfd7ccf596a5c8e80d56d))
* update versions to latest tag ([68ba5fb](https://github.com/devsy-org/api/commit/68ba5fb484c9d81f0e765fb65b797c965521b360))

## [3.0.0](https://github.com/devsy-org/api/compare/v2.0.0...v3.0.0) (2026-04-13)


### ⚠ BREAKING CHANGES

* API group renamed from loft.sh to devsy.sh. All Loft and VCluster type names renamed to Devsy equivalents.
* rename module from loft-sh/api/v4 to skevetter/api ([#11](https://github.com/devsy-org/api/issues/11))
* rename module from loft-sh/api/v4 to skevetter/api ([#10](https://github.com/devsy-org/api/issues/10))

### Features

* add annotation to exclude clusters for cluster account templates ([4a13959](https://github.com/devsy-org/api/commit/4a139590f6a067930422ecf0becd81c05dd510bc))
* add cluster account template frontend ([087dc69](https://github.com/devsy-org/api/commit/087dc69829e1c089ddad2f94f7880fd4812d6a15))
* add cluster reset for rbac & kiosk ([4e693a4](https://github.com/devsy-org/api/commit/4e693a4d5ff9c5cf6b97d03eb7f63277b8c8ccab))
* add profile picture ([1398c6a](https://github.com/devsy-org/api/commit/1398c6a03d20b96a2cb3f24191b6edc7b69f31b4))
* add sleepmode.loft.sh/delete-after ([0a96da7](https://github.com/devsy-org/api/commit/0a96da7dfe6e4d734648c095fbbbd497b291e6ae))
* add spaces sub view ([6e9c2ec](https://github.com/devsy-org/api/commit/6e9c2ec8505d4c4d0dd05d6c55a8940f2528dba7))
* add subdomain backend ([7651a1a](https://github.com/devsy-org/api/commit/7651a1a3fbd74537948e2d224db6613a7d13c6a0))
* chart username & password default ([7a4d167](https://github.com/devsy-org/api/commit/7a4d16752c8c096dcd00c13822843347f782bc88))
* cluster select frontend & antdesign update ([2844fbc](https://github.com/devsy-org/api/commit/2844fbcb94d35b8c64ab8956546adccad3aca9be))
* disable / enable users & token generation ([8bab235](https://github.com/devsy-org/api/commit/8bab235ad61f3e39ea2b2de4912ac83a1c2cc7b8))
* force sleep duration ([c4b5137](https://github.com/devsy-org/api/commit/c4b5137ac72123e17df2fc9236280c184b950962))
* image pull secrets ui ([e430a1c](https://github.com/devsy-org/api/commit/e430a1cb7b07a5098241a16cb956ff9c666342c6))
* improve loft login ([df87e59](https://github.com/devsy-org/api/commit/df87e59443f94c16a21035fd24fc36ff60c8248a))
* make air-gapped installation easier ([950171b](https://github.com/devsy-org/api/commit/950171bcdaf9a885eaa7b50cb9d404b1712ef0e3))
* **oidc:** add auth.oidc.insecureCa & auth.oidc.groups ([0863db8](https://github.com/devsy-org/api/commit/0863db83f87b64d29b03d333abdf046d63d602e3))
* **oidc:** implement support for thin oidc tokens ([8849a51](https://github.com/devsy-org/api/commit/8849a51412f5a51533e4688cc6458701f7a93bf4))
* rename loft/vcluster to devsy ([#14](https://github.com/devsy-org/api/issues/14)) ([c42a044](https://github.com/devsy-org/api/commit/c42a044a69cb358d6bad6cc89b51c8ea0a0b7889))
* rename module from loft-sh/api/v4 to skevetter/api ([#10](https://github.com/devsy-org/api/issues/10)) ([3241f19](https://github.com/devsy-org/api/commit/3241f197a59e36c90bb2490cade2f21b73a09669))
* rename module from loft-sh/api/v4 to skevetter/api ([#11](https://github.com/devsy-org/api/issues/11)) ([e2b7388](https://github.com/devsy-org/api/commit/e2b73888d90b4137fd8383c328bafa99d90313c7))
* restart loft automatically on config change & config comments ([bc7866e](https://github.com/devsy-org/api/commit/bc7866e7b64c23f81af95fb0a9eae5a67722b641))
* secrets frontend & crd ([c5e8b94](https://github.com/devsy-org/api/commit/c5e8b94adc6a1a64aff40c0ac66b1bcd133a524b))
* skip insecure tls verify helm & helm update ([1da9f07](https://github.com/devsy-org/api/commit/1da9f07129b2b6f945849f6e3567e0b470209100))
* sync cluster account templates ([cc416f2](https://github.com/devsy-org/api/commit/cc416f27b673c6800e04a7470b46ddcb059f0467))


### Bug Fixes

* add MegaLinter config and upgrade to v9 ([#17](https://github.com/devsy-org/api/issues/17)) ([17ddaaf](https://github.com/devsy-org/api/commit/17ddaaf08f767bbcb1a2d60298c2a2d818fd0944))
* fixes an issue where the profile picture is changed all the time ([846a0a1](https://github.com/devsy-org/api/commit/846a0a19a31f2939602352ee82b4fe4e6de2c3b7))
* fixes an issue with user & team api migration ([1b38d0c](https://github.com/devsy-org/api/commit/1b38d0c675cbf42d3f070e0c4769ef9ec899061f))
* fixes e2e tests ([7ddeec8](https://github.com/devsy-org/api/commit/7ddeec8d54f5e37a5aa6fbcece0eb44d26894cd7))
* remove non existent secret client ([992039e](https://github.com/devsy-org/api/commit/992039e2b9bd0265926a54ae8fdce0422c0e4a96))
* set release-please manifest to v1.0.0 ([#15](https://github.com/devsy-org/api/issues/15)) ([c16f14b](https://github.com/devsy-org/api/commit/c16f14b307a8fd8433d3ecfd4772456c2f2cfac6))
* update versions to latest tag ([64145fb](https://github.com/devsy-org/api/commit/64145fb02515d386eaebfd7ccf596a5c8e80d56d))
* update versions to latest tag ([68ba5fb](https://github.com/devsy-org/api/commit/68ba5fb484c9d81f0e765fb65b797c965521b360))

## [2.0.0](https://github.com/devsy-org/api/compare/v1.0.0...v2.0.0) (2026-04-13)


### ⚠ BREAKING CHANGES

* API group renamed from loft.sh to devsy.sh. All Loft and VCluster type names renamed to Devsy equivalents.
* rename module from loft-sh/api/v4 to skevetter/api ([#11](https://github.com/devsy-org/api/issues/11))
* rename module from loft-sh/api/v4 to skevetter/api ([#10](https://github.com/devsy-org/api/issues/10))

### Features

* add annotation to exclude clusters for cluster account templates ([4a13959](https://github.com/devsy-org/api/commit/4a139590f6a067930422ecf0becd81c05dd510bc))
* add cluster account template frontend ([087dc69](https://github.com/devsy-org/api/commit/087dc69829e1c089ddad2f94f7880fd4812d6a15))
* add cluster reset for rbac & kiosk ([4e693a4](https://github.com/devsy-org/api/commit/4e693a4d5ff9c5cf6b97d03eb7f63277b8c8ccab))
* add profile picture ([1398c6a](https://github.com/devsy-org/api/commit/1398c6a03d20b96a2cb3f24191b6edc7b69f31b4))
* add sleepmode.loft.sh/delete-after ([0a96da7](https://github.com/devsy-org/api/commit/0a96da7dfe6e4d734648c095fbbbd497b291e6ae))
* add spaces sub view ([6e9c2ec](https://github.com/devsy-org/api/commit/6e9c2ec8505d4c4d0dd05d6c55a8940f2528dba7))
* add subdomain backend ([7651a1a](https://github.com/devsy-org/api/commit/7651a1a3fbd74537948e2d224db6613a7d13c6a0))
* chart username & password default ([7a4d167](https://github.com/devsy-org/api/commit/7a4d16752c8c096dcd00c13822843347f782bc88))
* cluster select frontend & antdesign update ([2844fbc](https://github.com/devsy-org/api/commit/2844fbcb94d35b8c64ab8956546adccad3aca9be))
* disable / enable users & token generation ([8bab235](https://github.com/devsy-org/api/commit/8bab235ad61f3e39ea2b2de4912ac83a1c2cc7b8))
* force sleep duration ([c4b5137](https://github.com/devsy-org/api/commit/c4b5137ac72123e17df2fc9236280c184b950962))
* image pull secrets ui ([e430a1c](https://github.com/devsy-org/api/commit/e430a1cb7b07a5098241a16cb956ff9c666342c6))
* improve loft login ([df87e59](https://github.com/devsy-org/api/commit/df87e59443f94c16a21035fd24fc36ff60c8248a))
* make air-gapped installation easier ([950171b](https://github.com/devsy-org/api/commit/950171bcdaf9a885eaa7b50cb9d404b1712ef0e3))
* **oidc:** add auth.oidc.insecureCa & auth.oidc.groups ([0863db8](https://github.com/devsy-org/api/commit/0863db83f87b64d29b03d333abdf046d63d602e3))
* **oidc:** implement support for thin oidc tokens ([8849a51](https://github.com/devsy-org/api/commit/8849a51412f5a51533e4688cc6458701f7a93bf4))
* rename loft/vcluster to devsy ([#14](https://github.com/devsy-org/api/issues/14)) ([c42a044](https://github.com/devsy-org/api/commit/c42a044a69cb358d6bad6cc89b51c8ea0a0b7889))
* rename module from loft-sh/api/v4 to skevetter/api ([#10](https://github.com/devsy-org/api/issues/10)) ([3241f19](https://github.com/devsy-org/api/commit/3241f197a59e36c90bb2490cade2f21b73a09669))
* rename module from loft-sh/api/v4 to skevetter/api ([#11](https://github.com/devsy-org/api/issues/11)) ([e2b7388](https://github.com/devsy-org/api/commit/e2b73888d90b4137fd8383c328bafa99d90313c7))
* restart loft automatically on config change & config comments ([bc7866e](https://github.com/devsy-org/api/commit/bc7866e7b64c23f81af95fb0a9eae5a67722b641))
* secrets frontend & crd ([c5e8b94](https://github.com/devsy-org/api/commit/c5e8b94adc6a1a64aff40c0ac66b1bcd133a524b))
* skip insecure tls verify helm & helm update ([1da9f07](https://github.com/devsy-org/api/commit/1da9f07129b2b6f945849f6e3567e0b470209100))
* sync cluster account templates ([cc416f2](https://github.com/devsy-org/api/commit/cc416f27b673c6800e04a7470b46ddcb059f0467))


### Bug Fixes

* add MegaLinter config and upgrade to v9 ([#17](https://github.com/devsy-org/api/issues/17)) ([17ddaaf](https://github.com/devsy-org/api/commit/17ddaaf08f767bbcb1a2d60298c2a2d818fd0944))
* fixes an issue where the profile picture is changed all the time ([846a0a1](https://github.com/devsy-org/api/commit/846a0a19a31f2939602352ee82b4fe4e6de2c3b7))
* fixes an issue with user & team api migration ([1b38d0c](https://github.com/devsy-org/api/commit/1b38d0c675cbf42d3f070e0c4769ef9ec899061f))
* fixes e2e tests ([7ddeec8](https://github.com/devsy-org/api/commit/7ddeec8d54f5e37a5aa6fbcece0eb44d26894cd7))
* remove non existent secret client ([992039e](https://github.com/devsy-org/api/commit/992039e2b9bd0265926a54ae8fdce0422c0e4a96))
* set release-please manifest to v1.0.0 ([#15](https://github.com/devsy-org/api/issues/15)) ([c16f14b](https://github.com/devsy-org/api/commit/c16f14b307a8fd8433d3ecfd4772456c2f2cfac6))
* update versions to latest tag ([64145fb](https://github.com/devsy-org/api/commit/64145fb02515d386eaebfd7ccf596a5c8e80d56d))
* update versions to latest tag ([68ba5fb](https://github.com/devsy-org/api/commit/68ba5fb484c9d81f0e765fb65b797c965521b360))
