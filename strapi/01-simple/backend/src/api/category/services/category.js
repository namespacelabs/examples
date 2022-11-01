// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

'use strict';

/**
 * category service.
 */

const { createCoreService } = require('@strapi/strapi').factories;

module.exports = createCoreService('api::category.category');
