package com.shmulik.license.controller

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("v1/organization/{organizationId}/license")
class LicenseController {
    @GetMapping
    fun get(): String = "Hello"
} 