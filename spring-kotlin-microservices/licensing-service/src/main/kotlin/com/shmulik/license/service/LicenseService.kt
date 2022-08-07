package com.shmulik.license.service

import com.shmulik.license.model.License
import org.springframework.stereotype.Service
import kotlin.random.Random

@Service
public class LicenseService {
    fun getLicense(licenseId: String, organizationId: String): License {
        return License(Random.nextInt(1000), licenseId, "Software Product", organizationId, "Ostock", "full")
    }
}