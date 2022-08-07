package com.shmulik.license.controller

import com.shmulik.license.model.License
import com.shmulik.license.service.LicenseService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("v1/organization/{organizationId}/license")
class LicenseController {
    @Autowired
    lateinit var licenseService: LicenseService

    @GetMapping("/{licenseId}")
    fun get(
        @PathVariable("organizationId") organizationId: String,
        @PathVariable("licenseId") licenseId: String
    ): ResponseEntity<License> = ResponseEntity.ok(licenseService.getLicense(licenseId, organizationId))
} 