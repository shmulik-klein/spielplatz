package com.shmulik.license.model

data class License(
    val id: Int,
    val licenseId: String,
    val description: String,
    val organizationId: String,
    val productName: String,
    val licenseType: String
)
